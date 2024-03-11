package pattern

import "fmt"

/*
	Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.

	Плюсы:
	- Упрощает добавление операций, работающих со сложными структурами объектов.
	- Объединяет родственные операции в одном классе.
	- Посетитель может накапливать состояние при обходе структуры элементов.

	Минусы:
	- Паттерн не оправдан, если иерархия элементов часто меняется.
	- Может привести к нарушению инкапсуляции элементов.
*/

type Visitor interface {
	VisitFoo(Foo)
	VisitBar(Bar)
}

type Element struct {}

func (e *Element) VisitFoo(foo Foo) {
	fmt.Printf("Element visited Foo\n")
}

func (e *Element) VisitBar(bar Bar) {
	fmt.Printf("Element visited Bar\n")
}

type Acceptor interface {
	Accept(Visitor)
}

type Foo struct{}

type Bar struct{}

func (f *Foo) Accept(v Visitor) {
	v.VisitFoo(*f)
}

func (b *Bar) Accept(v Visitor) {
	v.VisitBar(*b)
}

func main() {
	items := []Acceptor{&Foo{}, &Bar{}}

	elem := &Element{}

	for _, item := range items {
		item.Accept(elem)
	}
}