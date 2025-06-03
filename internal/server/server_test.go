package server

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"chat_gopher/internal/config"
)

func TestNewServer(t *testing.T) {
	cfg := &config.Config{AppPort: "8080"}
	s := NewServer(cfg, &sql.DB{})
	if s.cfg.AppPort != "8080" {
		t.Errorf("Ожидался порт 8080, получили %s", s.cfg.AppPort)
	}
}

func TestIndexHandler(t *testing.T) {
	cfg := &config.Config{AppPort: "8080"}
	s := NewServer(cfg, &sql.DB{})

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался код 200, получили %d", status)
	}
}
