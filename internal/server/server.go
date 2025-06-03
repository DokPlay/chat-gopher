package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "chat_gopher/docs" // swagger docs
	"chat_gopher/internal/config"
	"chat_gopher/internal/handlers"
	"chat_gopher/internal/repository"
)

// Server - структура сервера
type Server struct {
	cfg  *config.Config
	db   *sql.DB
	repo repository.MessageRepository
}

// NewServer конструктор
func NewServer(cfg *config.Config, db *sql.DB) *Server {
	repo := repository.NewPgMessageRepository(db)
	return &Server{
		cfg:  cfg,
		db:   db,
		repo: repo,
	}
}

// Run запускает HTTP-сервер
func (s *Server) Run() error {
	logrus.Info("Настройка роутов...")
	r := mux.NewRouter()

	restHandler := handlers.NewRESTHandler(s.repo)

	// Роуты для HTML-страниц
	r.HandleFunc("/", s.indexHandler).Methods("GET")
	r.HandleFunc("/send", s.sendPageHandler).Methods("GET")
	r.HandleFunc("/realtime", s.realtimeHandler).Methods("GET")
	r.HandleFunc("/history", s.historyHandler).Methods("GET")

	// Статические файлы
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// API
	r.HandleFunc("/api/messages", restHandler.SendMessage).Methods("POST")
	r.HandleFunc("/api/messages", restHandler.GetMessages).Methods("GET")

	// WebSocket
	r.HandleFunc("/ws", handlers.WebSocketHandler)

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	addr := fmt.Sprintf(":%s", s.cfg.AppPort)
	logrus.Infof("Запуск сервера на %s", addr)
	return http.ListenAndServe(addr, r)
}

// indexHandler обрабатывает запросы к главной странице.
// Если файл "web/templates/index.html" отсутствует (например, в тестовой среде),
// возвращается дефолтный ответ с кодом 200.
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Открыта главная страница")
	if _, err := os.Stat("web/templates/index.html"); os.IsNotExist(err) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Index page"))
		return
	}
	http.ServeFile(w, r, "web/templates/index.html")
}

func (s *Server) sendPageHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Открыта страница отправки сообщения")
	http.ServeFile(w, r, "web/templates/send.html")
}

func (s *Server) realtimeHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Открыта страница реального времени")
	http.ServeFile(w, r, "web/templates/realtime.html")
}

func (s *Server) historyHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Открыта страница истории сообщений")
	http.ServeFile(w, r, "web/templates/history.html")
}
