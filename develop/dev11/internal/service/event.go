package service

import (
	"context"
	"dev11/internal/model"
	"dev11/internal/repository"
	"dev11/internal/request"
	"time"
)

type EventService interface {
	CreateEvent(ctx context.Context, request request.CreateEventRequest) (model.Event, error)
	UpdateEvent(ctx context.Context, request request.UpdateEventRequest) (model.Event, error)
	DeleteEvent(ctx context.Context, eventID int) error
	GetDayEvents(ctx context.Context, day time.Time) ([]model.Event, error)
	GetWeekEvents(ctx context.Context, week time.Time) ([]model.Event, error)
	GetMonthEvents(ctx context.Context, month time.Time) ([]model.Event, error)
}

type eventServiceImpl struct {
	Repository repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *eventServiceImpl {
	return &eventServiceImpl{
		Repository: repo,
	}
}

func (s *eventServiceImpl) CreateEvent(ctx context.Context, request request.CreateEventRequest) (model.Event, error) {
	err := request.Validate()

	if err != nil {
		return model.Event{}, err
	}

	event := model.Event{
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
		Date:        request.Date,
	}

	id, err := s.Repository.CreateEvent(ctx, event)
	if err != nil {
		return model.Event{}, err
	}

	event.ID = id
	return event, nil
}

func (s *eventServiceImpl) UpdateEvent(ctx context.Context, request request.UpdateEventRequest) (model.Event, error) {
	err := request.Validate()

	if err != nil {
		return model.Event{}, err
	}

	event := model.Event{
		ID:          request.ID,
		Title:       request.Title,
		Description: request.Description,
		Date:        request.Date,
	}

	err = s.Repository.UpdateEvent(ctx, event)
	if err != nil {
		return model.Event{}, err
	}

	return event, nil
}

func (s *eventServiceImpl) DeleteEvent(ctx context.Context, eventID int) error {
	return s.DeleteEvent(ctx, eventID)
}

func (s *eventServiceImpl) GetDayEvents(ctx context.Context, day time.Time) ([]model.Event, error) {
	return s.Repository.GetDayEvents(ctx, day)
}

func (s *eventServiceImpl) GetWeekEvents(ctx context.Context, week time.Time) ([]model.Event, error) {
	return s.Repository.GetDayEvents(ctx, week)
}

func (s *eventServiceImpl) GetMonthEvents(ctx context.Context, month time.Time) ([]model.Event, error) {
	return s.Repository.GetMonthEvents(ctx, month)
}