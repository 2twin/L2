package request

import (
	"errors"
	"time"
)

type CreateEventRequest struct {
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (r *CreateEventRequest) Validate() error {
	if r.UserID == 0 {
		return errors.New(`"user_id" field is required`)
	}

	if r.Title == "" {
		return errors.New(`"title" field is required`)
	}

	if r.Date.IsZero() {
		return errors.New(`"date" field is required`)
	}

	return nil
}

type UpdateEventRequest struct {
	ID          int
	Title       string
	Description string
	Date        time.Time
}

func (r *UpdateEventRequest) Validate() error {
	if r.ID == 0 {
		return errors.New(`"id" field is required`)
	}

	return nil
}
