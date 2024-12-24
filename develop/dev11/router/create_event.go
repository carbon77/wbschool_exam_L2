package router

import (
	"net/http"
	"ru/zakat/server/events"
)

func createEvent(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var event events.Event
		readBody(req, &event)

		events.Repository().AddEvent(&event)
		sendResultResponse(w, event)
	}
}
