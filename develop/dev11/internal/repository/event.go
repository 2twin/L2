package repository

import (
	"context"
	"dev11/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event model.Event) (int, error)
	GetEventByID(ctx context.Context, eventID int) (model.Event, error)
	UpdateEvent(ctx context.Context, event model.Event) error
	DeleteEvent(ctx context.Context, eventID int) error
	GetDayEvents(ctx context.Context, day time.Time) ([]model.Event, error)
	GetWeekEvents(ctx context.Context, week time.Time) ([]model.Event, error)
	GetMonthEvents(ctx context.Context, month time.Time) ([]model.Event, error)
}

type eventRepositoryImpl struct {
	DB *pgxpool.Pool
}

func NewEventRepository(pool *pgxpool.Pool) *eventRepositoryImpl {
	return &eventRepositoryImpl{
		DB: pool,
	}
}

func (r *eventRepositoryImpl) CreateEvent(ctx context.Context, event model.Event) (int, error) {
	var id int

	err := r.DB.QueryRow(ctx, `
		INSERT INTO events (user_id, title, description, date)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`, event.UserID, event.Title, event.Description, event.Date).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *eventRepositoryImpl) GetEventByID(ctx context.Context, eventID int) (model.Event, error) {
	var event model.Event

	err := r.DB.QueryRow(ctx, `
		SELECT * FROM events 
		WHERE id = $1;
	`, eventID).Scan(&event.ID, &event.UserID, &event.Title, &event.Description)

	if err != nil {
		return model.Event{}, err
	}

	return event, nil
}

func (r *eventRepositoryImpl) UpdateEvent(ctx context.Context, event model.Event) error {
	_, err := r.DB.Exec(ctx, `
		UPDATE events
		SET title = $1, description = $2, date = $3
		WHERE id = $4;
	`, event.Title, event.Description, event.Date, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepositoryImpl) DeleteEvent(ctx context.Context, eventID int) error {
	_, err := r.DB.Exec(ctx, `
		DELETE FROM events
		WHERE id = $1;
	`, eventID)

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepositoryImpl) GetDayEvents(ctx context.Context, day time.Time) ([]model.Event, error) {
	return []model.Event{}, nil
}

func (r *eventRepositoryImpl) GetWeekEvents(ctx context.Context, day time.Time) ([]model.Event, error) {
	return []model.Event{}, nil
}

func (r *eventRepositoryImpl) GetMonthEvents(ctx context.Context, day time.Time) ([]model.Event, error) {
	return []model.Event{}, nil
}
