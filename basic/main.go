package main

import "fmt"

/* definitions / behaviour / contract */

type Colour string

const (
	Red  Colour = "red"
	Blue Colour = "blue"
)

type PersonName interface {
	GetName() string
}

type Person interface {
	GetName() string
	GetAge() uint8
	ChangeHairColour(myColour Colour)
}

/* implementation */

type MyPerson struct {
	Name string
	Age  int
}

func (m *MyPerson) GetName() string {
	return m.Name
}

func (m *MyPerson) GetAge() uint8 {
	return uint8(m.Age)
}

func (m *MyPerson) ChangeHairColour(colour Colour) {
	// do something
}

/* code : using all this */

func main() {
	var myPerson = &MyPerson{
		Name: "John",
		Age:  21,
	}

	PrintPerson(myPerson)
	PrintNameOnly(myPerson)
}

func PrintPerson(p Person) {
	fmt.Printf("Name: %s, Age: %d\n", p.GetName(), p.GetAge())
}

func PrintNameOnly(p PersonName) {
	fmt.Printf("Name: %s\n", p.GetName())
}
