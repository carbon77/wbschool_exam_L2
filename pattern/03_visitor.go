package pattern

import (
	"fmt"
	"math"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Элемент (Element) описывает метод принятия посетителя
type Shape interface {
	Accept(visitor Visitor)
}

// Конкретный элемент (ConcreteElement) реализует интерфейс элемента для описания круга
type Circle struct {
	radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

// Конкретный элемент (ConcreteElement) реализует интерфейс элемента для описания квадрата
type Square struct {
	side float64
}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}

// Интерфейс посетителя, имеет два метода для работы с кругами и квадратами
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitSquare(square *Square)
}

// Конкретный посетитель (ConcreteVisitor) для вычисления площади фигур
type AreaVisitor struct{}

func (av AreaVisitor) VisitCircle(circle *Circle) {
	fmt.Printf("Площадь круга: %.2f\n", math.Pi*circle.radius*circle.radius)
}

func (av AreaVisitor) VisitSquare(square *Square) {
	fmt.Printf("Площадь квадрата: %.2f\n", square.side*square.side)
}

// Конкретный посетитель (ConcreteVisitor) для вычисления периметра фигур
type PerimeterVisitor struct{}

func (pv PerimeterVisitor) VisitCircle(circle *Circle) {
	fmt.Printf("Периметр круга: %.2f\n", 2*math.Pi*circle.radius)
}

func (pv PerimeterVisitor) VisitSquare(square *Square) {
	fmt.Printf("Периметр квадрата: %.2f\n", 4*square.side)
}

// Конкретный посетитель (ConcreteVisitor) для экспорта фигуры в формат XML
type XmlExportVisitor struct{}

func (xev XmlExportVisitor) VisitCircle(circle *Circle) {
	fmt.Printf("<Circle radius=%.2f />\n", circle.radius)
}

func (xev XmlExportVisitor) VisitSquare(square *Square) {
	fmt.Printf("<Square side=%.2f />\n", square.side)
}

// Структура клиента. Хранит коллекцию фигур и принимает посетителя для передачи его фигурам
type Client struct {
	shapes []Shape
}

func (c *Client) AddShape(shape Shape) {
	c.shapes = append(c.shapes, shape)
}

func (c *Client) Accept(visitor Visitor) {
	for _, shape := range c.shapes {
		shape.Accept(visitor)
	}
}

func TestVisitor() {
	shapes := []Shape{
		&Circle{3.5},
		&Square{5.6},
	}
	client := &Client{shapes}
	areaVisitor := &AreaVisitor{}
	perimeterVisitor := &PerimeterVisitor{}
	xmlExportVisitor := &XmlExportVisitor{}

	client.Accept(areaVisitor)
	client.Accept(perimeterVisitor)
	client.Accept(xmlExportVisitor)
}
