// @title My Messenger API
// @version 1.0
// @description Это пример сервиса обмена сообщениями
// @BasePath /

package main

import (
	"os"

	"chat_gopher/internal/config"
	"chat_gopher/internal/database"
	"chat_gopher/internal/server"

	"github.com/sirupsen/logrus"
)

// @title Chat Gopher API
// @version 1.0
// @description Это приложение для обмена сообщениями.
// @BasePath /
//
// @contact.name Developer
// @contact.url https://example.com
// @contact.email developer@example.com
//
// @host localhost:8080
// @schemes http

// main - точка входа в приложение.
func main() {
	// Настраиваем логирование с полными метками времени.
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.Info("Запуск приложения")

	// Загружаем конфигурацию.
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Ошибка загрузки конфигурации")
	}

	// Инициализируем подключение к базе данных.
	db, err := database.InitDB(cfg)
	if err != nil {
		logrus.WithError(err).Fatal("Ошибка инициализации БД")
	}

	// Создаем и запускаем сервер.
	srv := server.NewServer(cfg, db)
	if err := srv.Run(); err != nil {
		logrus.WithError(err).Fatal("Ошибка запуска сервера")
		os.Exit(1)
	}
}
