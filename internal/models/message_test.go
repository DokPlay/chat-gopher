package models

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestMessageFields(t *testing.T) {
	msg := Message{
		ID:        1,
		Text:      "Hello",
		Sequence:  100,
		CreatedAt: time.Now(),
	}

	// Использование поля CreatedAt для логирования
	logrus.Infof("Message created at: %s", msg.CreatedAt)

	// Проверка значений полей
	if msg.ID != 1 {
		t.Errorf("Expected ID 1, got %d", msg.ID)
	}
	if msg.Text != "Hello" {
		t.Errorf("Expected Text 'Hello', got '%s'", msg.Text)
	}
	if msg.Sequence != 100 {
		t.Errorf("Expected Sequence 100, got %d", msg.Sequence)
	}

	// Проверка, что CreatedAt не пустое
	if msg.CreatedAt.IsZero() {
		t.Errorf("Expected CreatedAt to be set, but got zero value")
	}
}
