package basic

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

// Compare this snippet from basic/main.go:

func main() {

	var myHero SuperHero = NewSuperHero("John", "Superman", 21)
	myHero.GetAge()

	myHero.GetName()
}

func (m MySuperPerson) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (m MySuperPerson) GetAge() int {
	//TODO implement me
	panic("implement me")
}

func (m MySuperPerson) GetHeroName() string {
	//TODO implement me
	panic("implement me")
}

func (m MySuperPerson) CanFly() bool {
	//TODO implement me
	panic("implement me")
}

func (m MySuperPerson) IsHuman() bool {
	//TODO implement me
	panic("implement me")
}

func (m MySuperPerson) IsSuperSpeed() bool {
	//TODO implement me
	panic("implement me")
}

func (m MySuperPerson) IsSuperStrength() bool {
	//TODO implement me
	panic("implement me")
}
