package router

import (
	"net/http"
	"ru/zakat/server/events"
)

type DeleteEventRequest struct {
	EventId string `json:"event_id"`
}

func deleteEvent(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var deleteRequest DeleteEventRequest
		err := readBody(req, &deleteRequest)
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		events.Repository().DeleteEvent(deleteRequest.EventId)
	}
}
