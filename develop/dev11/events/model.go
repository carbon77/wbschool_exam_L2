package events

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserId uint

type Event struct {
	Id       string    `json:"event_id"`
	UserId   UserId    `json:"user_id"`
	Title    string    `json:"title"`
	DateTime time.Time `json:"datetime"`
}

func (e Event) String() string {
	return fmt.Sprintf("Event[user_id=%d, title=%s, datetime=%v]", e.UserId, e.Title, e.DateTime)
}

// Структура Event реализовывает интерфейс json.Unmarshaler
// для более гибкого преобразования json в структуру
func (e *Event) UnmarshalJSON(data []byte) error {
	_event := &struct {
		Id       string `json:"event_id"`
		UserId   UserId `json:"user_id"`
		Title    string `json:"title"`
		DateTime string `json:"datetime"`
	}{}

	err := json.Unmarshal(data, &_event)
	if err != nil {
		return err
	}

	e.Id = _event.Id
	e.UserId = _event.UserId
	e.Title = _event.Title

	datetime, err := time.Parse("2006-01-02", _event.DateTime)
	if err != nil {
		return err
	}
	e.DateTime = datetime
	return nil
}
