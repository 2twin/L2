package pattern

import "fmt"

/*
	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов.

	Плюсы:
	- Избавляет класс от привязки к конкретным классам продуктов.
	- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	- Упрощает добавление новых продуктов в программу.
	- Реализует принцип открытости/закрытости.

	Минусы:
	- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.
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