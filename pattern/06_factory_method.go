package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Product - базовый тип создаваемого объекта
type Product interface {
	GetName() string
}

// Конкретные продукты, которые реализуют интерфейс Product
type ProductA struct{}

func (p *ProductA) GetName() string {
	return "Product A"
}

type ProductB struct{}

func (p *ProductB) GetName() string {
	return "Product B"
}

// Creator - тип, создающий продукты
type Creator interface {
	Create() *Product
}


// Конкретные создатели, реализующие интерфейс Creator
type CreatorA struct{}

func (c *CreatorA) Create() Product {
	return &ProductA{}
}

type CreatorB struct{}

func (c *CreatorB) Create() Product {
	return &ProductA{}
}

func main() {
	productACreator := &CreatorA{}
	productBCreator := &CreatorB{}

	productA := productACreator.Create()
	productB := productBCreator.Create()

	fmt.Println(productA.GetName())
	fmt.Println(productB.GetName())
}