package factory

import "go-interfaces-demo/businesscase/domain"

type MyAbbrPerson struct {
	Name       string
	Age        int
	IsDoctor   bool
	IsKnighted bool
}

func NewAbbrPerson(name string, age int, isDoctor bool, isKnighted bool) domain.Person {
	return &MyAbbrPerson{
		Name:       name,
		Age:        age,
		IsDoctor:   isDoctor,
		IsKnighted: isKnighted,
	}
}

func (m MyAbbrPerson) GetName() string {
	if m.IsDoctor {
		return "Dr. " + m.Name
	}
	if m.IsKnighted {
		return "Sir " + m.Name
	}
	return m.Name
}

func (m MyAbbrPerson) NeedsToShower() bool {
	if m.IsDoctor {
		return true
	}
	return false
}

func (m MyAbbrPerson) GetAge() int {
	return m.Age
}
