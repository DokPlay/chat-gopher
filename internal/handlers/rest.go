package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"chat_gopher/internal/models"
	"chat_gopher/internal/repository"

	"github.com/sirupsen/logrus"
)

// RESTHandler - хендлер для REST-запросов
type RESTHandler struct {
	repo repository.MessageRepository
}

// NewRESTHandler - конструктор
func NewRESTHandler(repo repository.MessageRepository) *RESTHandler {
	return &RESTHandler{repo: repo}
}

// SendMessageRequest - структура для чтения тела запроса на отправку
type SendMessageRequest struct {
	Text     string `json:"text"`
	Sequence int    `json:"sequence"`
}

// ErrorResponse - структура ошибки в JSON
type ErrorResponse struct {
	Error string `json:"error"`
}

// SendMessage godoc
// @Summary Отправить одно сообщение
// @Description Создает новое сообщение, записывает в БД и рассылает его через WebSocket
// @Tags messages
// @Accept  json
// @Produce  json
// @Param input body SendMessageRequest true "Тело запроса"
// @Success 200 {object} models.Message
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/messages [post]
func (h *RESTHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Обработка запроса SendMessage")
	var req SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.WithError(err).Error("Ошибка декодирования JSON")
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	if len(req.Text) > 128 {
		logrus.Error("Текст сообщения превышает 128 символов")
		http.Error(w, "Text too long", http.StatusBadRequest)
		return
	}

	msg := models.Message{
		Text:      req.Text,
		Sequence:  req.Sequence,
		CreatedAt: time.Now(),
	}
	ctx := context.Background()
	id, err := h.repo.InsertMessage(ctx, msg)
	if err != nil {
		logrus.WithError(err).Error("Ошибка сохранения сообщения")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	msg.ID = id

	// Отправим сообщение в WebSocket-поток (см. ws.go)
	BroadcastToClients(msg)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

// GetMessages godoc
// @Summary Получить список сообщений за диапазон дат
// @Description Возвращает список сообщений из БД за указанный период
// @Tags messages
// @Accept  json
// @Produce  json
// @Param from query string true "Начальная дата (RFC3339)"
// @Param to query string true "Конечная дата (RFC3339)"
// @Success 200 {array} models.Message
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/messages [get]
func (h *RESTHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Обработка запроса GetMessages")

	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	if fromStr == "" || toStr == "" {
		logrus.Warn("Не указаны from или to")
		http.Error(w, "from and to query params are required", http.StatusBadRequest)
		return
	}

	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		logrus.WithError(err).Error("Ошибка парсинга from")
		http.Error(w, "invalid from format", http.StatusBadRequest)
		return
	}
	to, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		logrus.WithError(err).Error("Ошибка парсинга to")
		http.Error(w, "invalid to format", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	messages, err := h.repo.GetMessagesByDateRange(ctx, from, to)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при получении сообщений из БД")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
