package pattern

import "fmt"

/*
	Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
	по цепочке обработчиков. Каждый последующий обработчик решает, обработать запрос или передавать дальше по цепи.

	Плюсы:
	- Уменьшает зависимость между клиентом и обработчиками.
	- Реализует принцип единственной обязанности.
	- Реализует принцип открытости/закрытости.

	Минусы:
	- Запрос может остаться никем не обработанным.й обработчик решает, может ли он о
*/

type Department interface {
	Execute(patient *Patient)
	SetNext(department Department)
}

type Patient struct {
	name           string
	hasMedicalBook bool
	hasSymptoms    bool
	hasPayment     bool
}

type Reception struct {
	next Department
}

func (reception *Reception) Execute(patient *Patient) {
	if !patient.hasMedicalBook {
		fmt.Printf("[reception] patient %s has no medical book\n", patient.name)
		return
	}

	fmt.Printf("[reception] patient %s successfully registered\n", patient.name)
	reception.next.Execute(patient)
}

func (reception *Reception) SetNext(department Department) {
	reception.next = department
}

type Doctor struct {
	next Department
}

func (doctor *Doctor) Execute(patient *Patient) {
	if !patient.hasSymptoms {
		fmt.Printf("[doctor] patient %s has no symptoms\n", patient.name)
		return
	}

	fmt.Printf("[doctor] patient %s successfully prescribed medicine\n", patient.name)
	doctor.next.Execute(patient)
}

func (doctor *Doctor) SetNext(department Department) {
	doctor.next = department
}

type Pharmacy struct {
	next Department
}

func (pharmacy *Pharmacy) Execute(patient *Patient) {
	if !patient.hasPayment {
		fmt.Printf("[pharmacy] patient %s has no payment\n", patient.name)
		return
	}

	fmt.Printf("[pharmacy] patient %s successfully paid\n", patient.name)
}

func (pharmacy *Pharmacy) SetNext(department Department) {
	pharmacy.next = department
}

func main() {
	patient := &Patient{
		name:           "John Doe",
		hasMedicalBook: true,
		hasSymptoms:    true,
		hasPayment:     false,
	}

	reception := &Reception{}
	doctor := &Doctor{}
	pharmacy := &Pharmacy{}

	reception.SetNext(doctor)
	doctor.SetNext(pharmacy)

	reception.Execute(patient)
}
