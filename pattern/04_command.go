package pattern

import "fmt"

/*
	Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.

	Плюсы:
	- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	- Позволяет реализовать простую отмену и повтор операций.
	- Позволяет реализовать отложенный запуск операций.
	- Позволяет собирать сложные команды из простых.
	- Реализует принцип открытости/закрытости.

	Минусы:
	- Усложняет код программы из-за введения множества дополнительных классов.
*/

// Базовый тип с набором методов
type DatabaseClient interface {
	Insert()
	Update()
	Select()
	Delete()
}

// Имплементация интерфейса DatabaseClient
type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) Insert() {
	fmt.Println("Inserting record...")
}

func (db *Database) Update() {
	fmt.Println("Updating record...")
}

func (db *Database) Delete() {
	fmt.Println("Deleting record...")
}

func (db *Database) Select() {
	fmt.Println("Selecting record...")
}

// Command - интерфейс комманды
type Command interface {
	Execute()
}

// Конкретные команды для выполнения определенных методов
type InsertCommand struct {
	Database *Database
}

type UpdateCommand struct {
	Database *Database
}

type SelectCommand struct {
	Database *Database
}

type DeleteCommand struct {
	Database *Database
}

// Имплементация интерфейса Command для каждой из команд
func (c *InsertCommand) Execute() {
	c.Database.Insert()
}

func (c *UpdateCommand) Execute() {
	c.Database.Update()
}

func (c *SelectCommand) Execute() {
	c.Database.Select()
}

func (c *DeleteCommand) Execute() {
	c.Database.Delete()
}

// Тип - разработчик, который будет вызывать команды
type developer struct {
	Insert Command
	Update Command
	Delete Command
	Select Command
}

type Developer interface {
	InsertRecord()
	UpdateRecord()
	DeleteRecord()
	SelectRecord()
}

func (d *developer) InsertRecord() {
	d.Insert.Execute()
}

func (d *developer) UpdateRecord() {
	d.Update.Execute()
}

func (d *developer) DeleteRecord() {
	d.Delete.Execute()
}

func (d *developer) SelectRecord() {
	d.Select.Execute()
}

// Конструктор разработчика
func NewDeveloper(ins, upd, del, sel Command) Developer {
	return &developer{
		Insert: ins,
		Update: upd,
		Delete: del,
		Select: sel,
	}
}

func main() {
	db := NewDatabase()
	dev := NewDeveloper(
		&InsertCommand{db},
		&UpdateCommand{db},
		&DeleteCommand{db},
		&SelectCommand{db},
	)

	dev.InsertRecord()
	dev.UpdateRecord()
	dev.DeleteRecord()
	dev.SelectRecord()
}
