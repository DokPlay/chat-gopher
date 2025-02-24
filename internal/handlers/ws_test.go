package handlers

import (
	"testing"

	"chat_gopher/internal/models"
)

// Тестируем только логику BroadcastToClients.
// Полностью end-to-end WebSocket тест обычно требует спец. окружения.
func TestBroadcastToClients(t *testing.T) {
	// Изначально нет клиентов
	if len(clients) != 0 {
		t.Errorf("Ожидалось 0 клиентов, получили %d", len(clients))
	}

	// Попробуем послать сообщение, проверим, что не упадём
	msg := models.Message{ID: 999, Text: "Broadcast Test"}
	BroadcastToClients(msg)
	// Если не упало, значит все ок (пока без реального соединения)
}
