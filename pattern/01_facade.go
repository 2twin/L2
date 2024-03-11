package pattern

import "fmt"

/*
Фасад — это структурный шаблон проектирования, который предоставляет простой интерфейс к сложной системе классов,
библиотеке или фреймворку.

Плюсы:
- Изолирует клиентов от компонентов сложной подсистемы.

Минусы:
- Фасад рискует стать слишком большим.
*/


// Реализация паттерна:

// какая-то сложная логика...
type Shape interface {
	Draw()
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("drawing a rectangle")
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("drawing a circle")
}

type Triangle struct{}

func (t *Triangle) Draw() {
	fmt.Println("drawing a triangle")
}


// shapeMaker - фасад
type shapeMaker struct {
	Circle    Shape
	Rectangle Shape
	Triangle  Shape
}

type ShapeMaker interface {
	DrawShape()
}

func (s *shapeMaker) DrawShape() {
	s.Circle.Draw()
	s.Rectangle.Draw()
	s.Triangle.Draw()
}

// конструктор фасада
func NewShapeMaker() ShapeMaker {
	return &shapeMaker{
		Circle:    &Circle{},
		Rectangle: &Rectangle{},
		Triangle:  &Triangle{},
	}
}

func main() {
	facade := NewShapeMaker()
	facade.DrawShape()
}
