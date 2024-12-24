package events

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type EventsRepository interface {
	AddEvent(event *Event) *Event
	DeleteEvent(eventId string)
	UpdateEvent(event *Event) error

	GetEvents() []*Event
	GetEventsByUserId(userId UserId) []*Event
	GetEventsBetweenDates(start, end time.Time) []*Event
}

type eventsRepository struct {
	mu     *sync.Mutex
	events map[string]*Event
}

var (
	er = &eventsRepository{
		events: make(map[string]*Event),
		mu:     &sync.Mutex{},
	}
)

func Repository() EventsRepository {
	return er
}

func (ep *eventsRepository) AddEvent(event *Event) *Event {
	ep.mu.Lock()
	event.Id = uuid.New().String()
	ep.events[event.Id] = event
	ep.mu.Unlock()
	return event
}

func (ep *eventsRepository) DeleteEvent(eventId string) {
	ep.mu.Lock()
	delete(ep.events, eventId)
	ep.mu.Unlock()
}

// Функция для получения событий по предикату
func (ep *eventsRepository) getEventsByPredicate(predicate func(*Event) bool) []*Event {
	events := make([]*Event, 0)
	for _, event := range ep.events {
		if predicate(event) {
			events = append(events, event)
		}
	}
	return events
}

func (ep *eventsRepository) GetEvents() []*Event {
	events := make([]*Event, 0, len(ep.events))
	for _, event := range ep.events {
		events = append(events, event)
	}
	return events
}

func (ep *eventsRepository) GetEventsByUserId(userId UserId) []*Event {
	return ep.getEventsByPredicate(func(e *Event) bool { return userId == e.UserId })
}

func (ep *eventsRepository) GetEventsBetweenDates(start, end time.Time) []*Event {
	return ep.getEventsByPredicate(func(e *Event) bool {
		eventAfterStart := start.Equal(e.DateTime) || start.Before(e.DateTime)
		eventBeforeEnd := end.Equal(e.DateTime) || end.After(e.DateTime)
		return eventAfterStart && eventBeforeEnd
	})
}

func (ep *eventsRepository) UpdateEvent(newEvent *Event) error {
	oldEvent, ok := ep.events[newEvent.Id]
	if !ok {
		return errors.New(fmt.Sprintf("Event with id=%s not found", newEvent.Id))
	}

	oldEvent.Title = newEvent.Title
	oldEvent.DateTime = newEvent.DateTime
	oldEvent.UserId = newEvent.UserId
	return nil
}
