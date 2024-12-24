package utils

import "time"

// Функции для получения границ текущего дня, недели и месяца

func GetDayBorders() (time.Time, time.Time) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 0, 1).Add(-time.Nanosecond)
	return start, end
}

func GetWeekBorders() (time.Time, time.Time) {
	now := time.Now()
	weekday := now.Weekday()
	// В time.Weekday значение 0 это воскресенье, поэтому необходимо заменить на 7
	if weekday == 0 {
		weekday = 7
	}

	start := time.Date(now.Year(), now.Month(), now.Day()-int(weekday)+1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return start, end
}

func GetMonthBorders() (time.Time, time.Time) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return start, end
}
