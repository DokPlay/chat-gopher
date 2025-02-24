package models

import "time"

// Message - модель сообщения
type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Sequence  int       `json:"sequence"`
	CreatedAt time.Time `json:"created_at"`
}
