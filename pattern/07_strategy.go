package pattern

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type RouteStrategy interface {
	BuildRoute(A, B string, kilometers float64) (time time.Duration, price float64)
}

type CarStrategy struct {
	AvgCarSpeed    float64
	AvgPetrolPrice float64
}

func (cs *CarStrategy) BuildRoute(A, B string, kilometers float64) (time.Duration, float64) {
	hours := kilometers / cs.AvgCarSpeed
	return time.Duration(hours * float64(time.Hour)), cs.AvgPetrolPrice * hours
}

type WalkingStrategy struct {
	AvgWalkingSpeed float64
}

func (ws *WalkingStrategy) BuildRoute(A, B string, kilometers float64) (time.Duration, float64) {
	hours := kilometers / ws.AvgWalkingSpeed
	return time.Duration(hours * float64(time.Hour)), 0
}

type PublicTransportStrategy struct {
	AvgTrainSpeed float64
	AvgTrainPrice float64
}

func (pts *PublicTransportStrategy) BuildRoute(A, B string, kilometers float64) (time.Duration, float64) {
	hours := kilometers / pts.AvgTrainSpeed
	return time.Duration(hours * float64(time.Hour)), pts.AvgTrainPrice * hours
}

type Navigator struct {
	Strategies []RouteStrategy
}

func (n *Navigator) BuildRoute(A, B string, kilometers float64) {
	fmt.Printf("Путь между %s и %s:\n", A, B)
	for _, routeStrategy := range n.Strategies {
		travelTime, price := routeStrategy.BuildRoute(A, B, kilometers)

		fmt.Print("\t")
		switch routeStrategy.(type) {
		case *CarStrategy:
			fmt.Print("На машине: ")
		case *WalkingStrategy:
			fmt.Print("Пешком: ")
		case *PublicTransportStrategy:
			fmt.Print("На общ. транспорте: ")
		}
		fmt.Printf("время = %.2f час, цена = %.2f руб\n", travelTime.Hours(), price)
	}
}

func TestStrategy() {
	carStrategy := &CarStrategy{100, 400}
	walkingStrategy := &WalkingStrategy{4}
	publicTransportStrategy := &PublicTransportStrategy{60, 200}
	navigator := &Navigator{
		Strategies: []RouteStrategy{carStrategy, walkingStrategy, publicTransportStrategy},
	}
	navigator.BuildRoute("Москва", "Зеленоград", 42)
	navigator.BuildRoute("Казань", "Москва", 720)
}
