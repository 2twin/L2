package pattern

import "fmt"

/*
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
	каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

	Плюсы:
	- Горячая замена алгоритмов на лету.
	- Изолирует код и данные алгоритмов от остальных классов.
	- Уход от наследования к делегированию.
	- Реализует принцип открытости/закрытости.

	Минусы:
	- Усложняет программу за счёт дополнительных классов.
	- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

// базовый тип Human с полями: имя и инструмент
type Human struct {
	name string
	tool Tool
}

// методы Human
func (h *Human) Write() {
	h.tool.Write()
}

func (h *Human) SetTool(tool Tool) {
	h.tool = tool
}

// интерфейс инструемента с единственным методом Write - что-то записать
type Tool interface {
	Write()
}

// тип Pen - ручка
type Pen struct{}

// реализация интерфейса Tool для Pen
func (p *Pen) Write() {
	fmt.Println("Writing with pen")
}

// тип Brush - кисть
type Brush struct{}

// реализация интерфейса Tool для Pen
func (b *Brush) Write() {
	fmt.Println("Painting with brush")
}

// тип студент со встоенной структурой Human
type Student struct {
	h Human
}

// конструктор студента с инструментом Pen
func NewStudent(name string) *Student {
	return &Student{
		h: Human{
			name: name,
			tool: &Pen{},
		},
	}
}

// тип художник со встоенной структурой Human
type Painter struct {
	h Human
}

// конструктор студента с инструментом Brush
func NewPainter(name string) *Painter {
	return &Painter{
		h: Human{
			name: name,
			tool: &Brush{},
		},
	}
}

func main() {
	// создаем студента и вызываем метод Write
	pasha := NewStudent("Pasha")
	pasha.h.Write()

	// создаем художника и вызываем метод Write
	sasha := NewPainter("Sasha")
	sasha.h.Write()

	// теперь меняем инструмент у художника на Pen и вызываем Write
	sasha.h.SetTool(&Pen{})
	sasha.h.Write()
}
