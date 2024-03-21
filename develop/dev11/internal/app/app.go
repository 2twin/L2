package app

import (
	"dev11/internal/request"
	"dev11/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"time"
	// "time"
)

type App struct {
	service service.EventService
	addr    string
}

func NewApp(service service.EventService, addr string) *App {
	return &App{
		service: service,
		addr:    addr,
	}
}

func (a *App) Run() {
	http.HandleFunc("/create_event", a.HandlerCreateEvent)
	http.HandleFunc("/update_event", a.HandlerUpdateEvent)
	http.HandleFunc("/delete_event", a.HandlerDeleteEvent)
	http.HandleFunc("/events_for_day", a.HandlerGetEventsForDay)
	http.HandleFunc("/events_for_week", a.HandlerGetEventsForWeek)
	http.HandleFunc("/events_for_month", a.HandlerGetEventsForMonth)

	http.ListenAndServe(a.addr, LoggingMiddleware(http.DefaultServeMux))
}

func (a *App) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	request := request.CreateEventRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = request.Validate()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	event, err := a.service.CreateEvent(r.Context(), request)
	if err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, EventResponse{
		Result: event,
	})
}

func (a *App) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	request := request.UpdateEventRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = request.Validate()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	event, err := a.service.UpdateEvent(r.Context(), request)
	if err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, EventResponse{
		Result: event,
	})
}

func (a *App) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	type DeleteRequest struct {
		ID int `json:"id"`
	}
	request := DeleteRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = a.service.DeleteEvent(r.Context(), request.ID)
	if err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, EventResponse{
		Result: request,
	})
}

func (a *App) HandlerGetEventsForDay(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.URL.Query().Get("date"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	events, err := a.service.GetDayEvents(r.Context(), date)
	if err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, EventResponse{
		Result: events,
	})
}

func (a *App) HandlerGetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.URL.Query().Get("date"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	events, err := a.service.GetWeekEvents(r.Context(), date)
	if err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, EventResponse{
		Result: events,
	})
}

func (a *App) HandlerGetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.URL.Query().Get("date"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	events, err := a.service.GetMonthEvents(r.Context(), date)
	if err != nil {
		RespondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, EventResponse{
		Result: events,
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
