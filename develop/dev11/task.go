package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"log"
	"net/http"
)

// Domain object
type Event struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

// Response object
type Response struct {
	Result Event  `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Repository struct {
	Events map[int]Event
}

// Helper function to serialize object to JSON
func toJSON(obj interface{}) []byte {
	result, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
	}
	return result
}

// Helper function to handle HTTP errors
func handleError(w http.ResponseWriter, statusCode int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{Error: err}
	w.Write(toJSON(response))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
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

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// HTTP handler for /create_event
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your business logic here
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")

	if id == "" || name == "" || date == "" {
		// Example of handling an error
		handleError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusCreated, Response{
		Result: Event{
			ID:   id,
			Name: name,
			Date: date,
		},
	})
}

// HTTP handler for /update_event
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your business logic here
}

// HTTP handler for /delete_event
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your business logic here
}

// HTTP handler for /events_for_day
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your business logic here
}

// HTTP handler for /events_for_week
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your business logic here
}

// HTTP handler for /events_for_month
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your business logic here
}

func main() {
	// Setting up HTTP routes
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	// Applying logging middleware
	http.ListenAndServe(":8080", loggingMiddleware(http.DefaultServeMux))
}
