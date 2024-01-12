package main

import "fmt"

/* definitions / behaviour */

type Person interface {
	GetName() string
	GetAge() int
}

type SuperHero interface {
	Person
	GetHeroName() string
	CanFly() bool
	IsHuman() bool
	IsSuperSpeed() bool
	IsSuperStrength() bool
}

/* implementation */

type MySuperPerson struct {
	name     string
	heroName string
	age      int
}

func NewSuperHero(name string, heroName string, age int) SuperHero {
	return &MySuperPerson{
		name:     name,
		heroName: heroName,
		age:      age,
	}
}

func (m *MySuperPerson) GetName() string {
	return m.name
}

func (m *MySuperPerson) GetAge() int {
	return m.age
}

func (m *MySuperPerson) GetHeroName() string {
	return m.heroName
}

func (m *MySuperPerson) CanFly() bool {
	return false
}

func (m *MySuperPerson) IsHuman() bool {
	return true
}

func (m *MySuperPerson) IsSuperSpeed() bool {
	return true
}

func (m *MySuperPerson) IsSuperStrength() bool {
	return true
}

func main() {
	var myHero SuperHero = NewSuperHero("Bruce Wayne", "Batman", 45)

	fmt.Printf("My hero is %s\n", myHero.GetHeroName())
	fmt.Printf("His real name is %s\n", myHero.GetName())
	fmt.Printf("he is %d years old\n", myHero.GetAge())
}
