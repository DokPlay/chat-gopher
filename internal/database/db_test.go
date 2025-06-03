package database

import (
	"os"
	"testing"

	"chat_gopher/internal/config"
)

// Тест проверяет, что можем инициализировать in-memory (но для PostgreSQL это сложнее)
// Здесь, как вариант, просто тест на открытие/ошибку.
func TestInitDB(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "messages_db")

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Не удалось загрузить конфиг: %v", err)
	}

	db, err := InitDB(cfg)
	if err != nil {
		t.Fatalf("Не удалось инициализировать БД: %v", err)
	}
	if db == nil {
		t.Fatalf("Ожидалось соединение с БД, получили nil")
	}
	db.Close()
}
