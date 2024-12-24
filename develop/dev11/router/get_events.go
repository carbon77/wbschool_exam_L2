package router

import (
	"net/http"
	"ru/zakat/server/events"
	"ru/zakat/server/utils"
	"strconv"
)

func getEvents(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		userIdParam := getQueryParam(req, "user_id", "")
		var result []*events.Event

		if userIdParam == "" {
			result = events.Repository().GetEvents()
		} else {
			userId, err := strconv.Atoi(userIdParam)
			if err != nil {
				sendErrorResponse(w, "invalid user id", http.StatusBadRequest)
				return
			}
			result = events.Repository().GetEventsByUserId(events.UserId(userId))
		}

		sendResultResponse(w, result)
	}
}

func getEventsForDay(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		start, end := utils.GetDayBorders()
		result := events.Repository().GetEventsBetweenDates(start, end)
		sendResultResponse(w, result)
	}
}

func getEventsForWeek(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		start, end := utils.GetWeekBorders()
		result := events.Repository().GetEventsBetweenDates(start, end)
		sendResultResponse(w, result)
	}
}

func getEventsForMonth(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		start, end := utils.GetMonthBorders()
		result := events.Repository().GetEventsBetweenDates(start, end)
		sendResultResponse(w, result)
	}
}
