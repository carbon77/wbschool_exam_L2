package router

import (
	"net/http"
	"ru/zakat/server/events"
)

func updateEvent(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var event events.Event
		readBody(req, &event)

		err := events.Repository().UpdateEvent(&event)
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		sendResultResponse(w, event)
	}
}
