package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("DB_HOST", "test_host")
	os.Setenv("DB_PORT", "6543")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_PASSWORD", "test_pass")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("APP_PORT", "9999")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Ожидалось отсутствие ошибки, получили: %v", err)
	}

	if cfg.DBHost != "test_host" {
		t.Errorf("Ожидалось %s, получили %s", "test_host", cfg.DBHost)
	}
	if cfg.DBPort != 6543 {
		t.Errorf("Ожидалось 6543, получили %d", cfg.DBPort)
	}
	if cfg.DBUser != "test_user" {
		t.Errorf("Ожидалось test_user, получили %s", cfg.DBUser)
	}
	if cfg.DBPassword != "test_pass" {
		t.Errorf("Ожидалось test_pass, получили %s", cfg.DBPassword)
	}
	if cfg.DBName != "test_db" {
		t.Errorf("Ожидалось test_db, получили %s", cfg.DBName)
	}
	if cfg.AppPort != "9999" {
		t.Errorf("Ожидалось 9999, получили %s", cfg.AppPort)
	}
}
