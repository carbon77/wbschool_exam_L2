package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type MobileState int

const (
	POWER_OFF MobileState = iota
	SCREEN_DISABLED
	SCREEN_ON
)

type MobileScreen struct {
	state MobileState
}

func (ms *MobileScreen) PowerOn() {
	ms.state = SCREEN_DISABLED
}

func (ms *MobileScreen) Touch() {
	if ms.state == POWER_OFF {
		return
	}

	if ms.state == SCREEN_DISABLED {
		ms.state = SCREEN_ON
		return
	}

	fmt.Println("Touch!")
}

func (ms *MobileScreen) Swipe() {
	if ms.state != SCREEN_ON {
		return
	}

	fmt.Println("Swipe!")
}

func TestState() {
	mobileScreen := &MobileScreen{}
	mobileScreen.Touch()
	mobileScreen.PowerOn()
	mobileScreen.Touch()
	mobileScreen.Touch()
}
