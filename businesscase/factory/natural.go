package factory

import "go-interfaces-demo/businesscase/domain"

type MyPerson struct {
	Name string
	Age  int
}

func NewNaturalPerson(name string, age int) domain.Person {
	return &MyPerson{
		Name: name,
		Age:  age,
	}
}

func (m MyPerson) GetName() string {
	return m.Name
}

func (m MyPerson) GetAge() int {
	return m.Age
}

func (m MyPerson) NeedsToShower() bool {
	if m.Name == "John" {
		return true
	}
	return false
}
