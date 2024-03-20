package app

import (
	"dev11/internal/model"
	"dev11/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

const (
	DAY   = "day"
	WEEK  = "week"
	MONTH = "MONTH"
)

type App struct {
	service service.EventService
	addr    string
}

func NewApp(service service.EventService, addr string) *App {
	return &App{
		service: service,
		addr: addr,
	}
}

func (a *App) Run() {
	http.HandleFunc("/create_event", HandlerCreateEvent)

	http.ListenAndServe(a.addr, LoggingMiddleware(http.DefaultServeMux))
}

type EventResponse struct {
	Result model.Event `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	RespondWithJSON(w, code, ErrorResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {

}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
