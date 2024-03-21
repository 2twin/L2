package model

import "time"

type Event struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date"`
}