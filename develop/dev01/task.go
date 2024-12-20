package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/
func Now() time.Time {
	timeResult, err := ntp.Time("time.google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в получении времени: %v\n", err)
		os.Exit(1)
	}

	return timeResult
}

func main() {
	now := Now()
	fmt.Printf("Текущее время: %v\n", now.Format(time.UnixDate))
}
