package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Config структура для хранения конфигурации
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

// LoadConfig загружает конфигурацию из переменных окружения (или задаёт по умолчанию)
func LoadConfig() (*Config, error) {
	logrus.Info("Загрузка конфигурации...")

	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		logrus.WithError(err).Error("Невозможно конвертировать DB_PORT в int")
		return nil, err
	}

	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     port,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "messages_db"),
		AppPort:    getEnv("APP_PORT", "8080"),
	}

	logrus.Infof("Конфигурация загружена: %+v", cfg)
	return cfg, nil
}

func getEnv(key, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		logrus.Warnf("Переменная окружения %s не установлена, используем значение по умолчанию: %s", key, defaultVal)
		return defaultVal
	}
	return val
}
