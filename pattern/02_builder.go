package pattern

import "fmt"

/*
	Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
	Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

	Плюсы:
	- Позволяет создавать продукты пошагово.
	- Позволяет использовать один и тот же код для создания различных продуктов.
	- Изолирует сложный код сборки продукта от его основной бизнес-логики.

	Минусы:
	- Усложняет код программы из-за введения дополнительных классов.
	- Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

// Базовый объект строительства
type Pizza struct {
	Pastry    string
	Sauce     string
	Garniture string
}

// Builder - строитель
type PizzaBuilder interface {
	CreateNewPizza()
	GetPizza() string
	BuildPastry()
	BuildSauce()
	BuildGarniture()
}

// Конкретный строитель 1
type BuilderPizzaHawaii struct {
	pizza *Pizza
}

func NewBuilderPizzaHawaii() *BuilderPizzaHawaii {
	return &BuilderPizzaHawaii{}
}

func (b *BuilderPizzaHawaii) GetPizza() string {
	return fmt.Sprintf("Your order. pastry: %s, sauce: %s, garniture: %s", b.pizza.Pastry, b.pizza.Sauce, b.pizza.Garniture)
}

func (b *BuilderPizzaHawaii) CreateNewPizza() {
	b.pizza = &Pizza{}
}

func (b *BuilderPizzaHawaii) BuildPastry() {
	b.pizza.Pastry = "normal"
}

func (b *BuilderPizzaHawaii) BuildSauce() {
	b.pizza.Sauce = "soft"
}

func (b *BuilderPizzaHawaii) BuildGarniture() {
	b.pizza.Garniture = "jambon + ananas"
}

// Конкретный строитель 2
type BuilderPizzaSpicy struct {
	pizza *Pizza
}

func NewBuilderPizzaSpicy() *BuilderPizzaSpicy {
	return &BuilderPizzaSpicy{}
}

func (b *BuilderPizzaSpicy) GetPizza() string {
	return fmt.Sprintf("Your order. pastry: %s, sauce: %s, garniture: %s", b.pizza.Pastry, b.pizza.Sauce, b.pizza.Garniture)
}

func (b *BuilderPizzaSpicy) CreateNewPizza() {
	b.pizza = &Pizza{}
}

func (b *BuilderPizzaSpicy) BuildPastry() {
	b.pizza.Pastry = "puff"
}

func (b *BuilderPizzaSpicy) BuildSauce() {
	b.pizza.Sauce = "hot"
}

func (b *BuilderPizzaSpicy) BuildGarniture() {
	b.pizza.Garniture = "papperoni+salami"
}

// Director - Управляющий класс, запускающий строительство
type Waiter interface {
	SetBuilderPizza(PizzaBuilder)
	GetPizza() string
	ConstructPizza()
}

type waiter struct {
	Builder PizzaBuilder
}

func NewWaiter() Waiter {
	return &waiter{}
}

func (w *waiter) SetBuilderPizza(pb PizzaBuilder) {
	w.Builder = pb
}

func (w *waiter) GetPizza() string {
	return w.Builder.GetPizza()
}

func (w *waiter) ConstructPizza() {
	w.Builder.CreateNewPizza()
	w.Builder.BuildPastry()
	w.Builder.BuildSauce()
	w.Builder.BuildGarniture()
}

func main() {
	w := NewWaiter()

	builderPizzaHawaii := NewBuilderPizzaHawaii()
	builderPizzaSpicy := NewBuilderPizzaSpicy()

	w.SetBuilderPizza(builderPizzaHawaii)
	w.ConstructPizza()
	fmt.Println(w.GetPizza())

	w.SetBuilderPizza(builderPizzaSpicy)
	w.ConstructPizza()
	fmt.Println(w.GetPizza())
}
