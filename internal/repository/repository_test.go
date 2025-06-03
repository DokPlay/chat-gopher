package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"chat_gopher/internal/models"

	_ "github.com/lib/pq"
)

func TestPgMessageRepository(t *testing.T) {
	// В реальных условиях нужно мокать DB или поднимать тестовую БД.
	// Для примера же можно протестировать на реальной локальной БД (Docker).
	dsn := "host=localhost port=5432 user=postgres password=password dbname=messages_db sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	repo := NewPgMessageRepository(db)

	// Тест InsertMessage
	msg := models.Message{
		Text:      "Test",
		Sequence:  1,
		CreatedAt: time.Now(),
	}
	id, err := repo.InsertMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("Ошибка InsertMessage: %v", err)
	}
	if id == 0 {
		t.Errorf("Ожидался ID > 0, получили 0")
	}

	// Тест GetMessagesByDateRange
	from := time.Now().Add(-1 * time.Hour)
	to := time.Now().Add(1 * time.Hour)
	messages, err := repo.GetMessagesByDateRange(context.Background(), from, to)
	if err != nil {
		t.Fatalf("Ошибка GetMessagesByDateRange: %v", err)
	}
	if len(messages) == 0 {
		t.Errorf("Ожидалось хотя бы одно сообщение, получили 0")
	}
}
