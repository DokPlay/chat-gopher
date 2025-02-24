package database

import (
	"database/sql"
	"fmt"

	"chat_gopher/internal/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// InitDB инициализирует подключение к PostgreSQL
func InitDB(cfg *config.Config) (*sql.DB, error) {
	logrus.Info("Инициализация подключения к БД...")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при открытии подключения к БД")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.WithError(err).Error("Ошибка при проверке соединения с БД")
		return nil, err
	}

	logrus.Info("Подключение к БД успешно!")

	// Создадим таблицу messages, если ее нет
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		text VARCHAR(128) NOT NULL,
		sequence INT NOT NULL,
		created_at TIMESTAMP NOT NULL
	);
	`)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при создании таблицы messages")
		return nil, err
	}
	logrus.Info("Таблица messages проверена/создана")

	return db, nil
}
