package handlers

import (
	"net/http"
	"sync"

	"chat_gopher/internal/models"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// upgrader используется для апгрейда HTTP-соединения до WebSocket.
// CheckOrigin разрешает подключение от любых источников (для упрощения).
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// clientsMu защищает мапу active WebSocket-соединений.
var clientsMu sync.Mutex

// clients хранит активные WebSocket-соединения.
var clients = make(map[*websocket.Conn]bool)

// WebSocketHandler обрабатывает HTTP-запросы для установления WebSocket-соединения.
//
// @Summary Устанавливает WebSocket-соединение
// @Description Апгрейдит HTTP-соединение до WebSocket и регистрирует клиента.
// @Tags websocket
// @Accept json
// @Produce json
// @Success 101 {string} string "Switching Protocols"
// @Failure 500 {object} map[string]string "Internal Server Error"
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Обработка WebSocket-подключения")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithError(err).Error("Ошибка апгрейда соединения до WebSocket")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	logrus.Info("WebSocket-подключение успешно")

	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	// Запускаем чтение сообщений, чтобы при закрытии клиента удалить соединение.
	go readLoop(conn)
}

// readLoop читает сообщения из WebSocket-соединения в цикле.
// При возникновении ошибки чтения соединение закрывается и удаляется из списка клиентов.
func readLoop(conn *websocket.Conn) {
	defer func() {
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
		conn.Close()
		logrus.Info("WebSocket-подключение закрыто")
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			logrus.WithError(err).Warn("Ошибка чтения из WebSocket, закрываем соединение")
			break
		}
	}
}

// BroadcastToClients рассылает переданное сообщение всем подключенным WebSocket-клиентам.
//
// @Summary Рассылка сообщения всем клиентам
// @Description Отправляет JSON-представление сообщения каждому подключенному клиенту. При ошибке отправки соединение удаляется.
// @Tags websocket
// @Accept json
// @Produce json
func BroadcastToClients(msg models.Message) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for c := range clients {
		err := c.WriteJSON(msg)
		if err != nil {
			logrus.WithError(err).Warn("Ошибка при отправке сообщения клиенту, удаляем соединение")
			c.Close()
			delete(clients, c)
		}
	}
}
