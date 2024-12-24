package router

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResultResponse struct {
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var (
	// Обработчики методов
	handlers = map[string]http.HandlerFunc{
		"/create_event": createEvent,
		"/delete_event": deleteEvent,
		"/update_event": updateEvent,
		"/events":       getEvents,

		// Методы для получения событий на текущий день, неделю, месяц
		"/events_for_day":   getEventsForDay,
		"/events_for_week":  getEventsForWeek,
		"/events_for_month": getEventsForMonth,
	}

	middlewares = []Middleware{
		logMiddleware,
	}
)

// Инициализация обработчиков
func InitRouter() {
	for path, handleFunc := range handlers {
		handler := handleFunc
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		http.HandleFunc(path, handler)
	}
}

// Функция для отправки результата запроса
func sendResultResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ResultResponse{data})
}

// Функция для отправки
func sendErrorResponse(w http.ResponseWriter, message string, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(ErrorResponse{message})
}

// Функция для получения query параметров или значения по умолчанию
func getQueryParam(req *http.Request, paramName, defaultValue string) string {
	if req.URL.Query().Has(paramName) {
		return req.URL.Query().Get(paramName)
	}
	return defaultValue
}

// Функция для чтения тела запроса в json в объект obj
func readBody(req *http.Request, obj any) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, obj); err != nil {
		return err
	}
	return nil
}
