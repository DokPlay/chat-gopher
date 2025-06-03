package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"chat_gopher/internal/models"
)

type mockRepo struct{}

func (m *mockRepo) InsertMessage(ctx context.Context, msg models.Message) (int, error) {
	return 123, nil
}
func (m *mockRepo) GetMessagesByDateRange(ctx context.Context, from, to time.Time) ([]models.Message, error) {
	return []models.Message{{ID: 1, Text: "Mock", Sequence: 10, CreatedAt: time.Now()}}, nil
}

func TestSendMessage(t *testing.T) {
	handler := NewRESTHandler(&mockRepo{})
	reqBody := []byte(`{"text":"Hello","sequence":1}`)
	req, _ := http.NewRequest("POST", "/api/messages", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	handler.SendMessage(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидали статус 200, получили %d", status)
	}

	var msg models.Message
	json.Unmarshal(rr.Body.Bytes(), &msg)
	if msg.ID != 123 {
		t.Errorf("Ожидался ID=123, получили %d", msg.ID)
	}
	if msg.Text != "Hello" {
		t.Errorf("Ожидался текст 'Hello', получили '%s'", msg.Text)
	}
}

func TestGetMessages(t *testing.T) {
	handler := NewRESTHandler(&mockRepo{})
	req, _ := http.NewRequest("GET", "/api/messages?from=2021-01-01T00:00:00Z&to=2023-01-01T00:00:00Z", nil)
	rr := httptest.NewRecorder()

	handler.GetMessages(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидали статус 200, получили %d", status)
	}

	var messages []models.Message
	json.Unmarshal(rr.Body.Bytes(), &messages)
	if len(messages) == 0 {
		t.Errorf("Ожидался массив с 1 сообщением, получили 0")
	}
}
