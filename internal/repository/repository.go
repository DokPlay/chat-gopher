package repository

import (
	"context"
	"database/sql"
	"time"

	"chat_gopher/internal/models"

	"github.com/sirupsen/logrus"
)

// MessageRepository - интерфейс для работы с сообщениями
type MessageRepository interface {
	InsertMessage(ctx context.Context, msg models.Message) (int, error)
	GetMessagesByDateRange(ctx context.Context, from, to time.Time) ([]models.Message, error)
}

// PgMessageRepository - реализация MessageRepository для PostgreSQL
type PgMessageRepository struct {
	db *sql.DB
}

func NewPgMessageRepository(db *sql.DB) *PgMessageRepository {
	return &PgMessageRepository{
		db: db,
	}
}

// InsertMessage вставляет новое сообщение в БД
func (r *PgMessageRepository) InsertMessage(ctx context.Context, msg models.Message) (int, error) {
	logrus.WithFields(logrus.Fields{
		"text":     msg.Text,
		"sequence": msg.Sequence,
	}).Info("Вставка сообщения в БД")
	query := `INSERT INTO messages (text, sequence, created_at) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query, msg.Text, msg.Sequence, msg.CreatedAt).Scan(&id)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при вставке сообщения")
		return 0, err
	}
	logrus.Infof("Сообщение вставлено с id=%d", id)
	return id, nil
}

// GetMessagesByDateRange получает сообщения за указанный диапазон дат
func (r *PgMessageRepository) GetMessagesByDateRange(ctx context.Context, from, to time.Time) ([]models.Message, error) {
	logrus.WithFields(logrus.Fields{
		"from": from,
		"to":   to,
	}).Info("Получение сообщений по диапазону дат")
	query := `SELECT id, text, sequence, created_at FROM messages WHERE created_at BETWEEN $1 AND $2 ORDER BY created_at ASC`
	rows, err := r.db.QueryContext(ctx, query, from, to)
	if err != nil {
		logrus.WithError(err).Error("Ошибка при выполнении запроса на выборку сообщений")
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.Text, &m.Sequence, &m.CreatedAt); err != nil {
			logrus.WithError(err).Error("Ошибка при чтении строки из выборки")
			return nil, err
		}
		messages = append(messages, m)
	}
	if err = rows.Err(); err != nil {
		logrus.WithError(err).Error("Ошибка после чтения всех строк")
		return nil, err
	}
	logrus.Infof("Найдено %d сообщений", len(messages))
	return messages, nil
}
