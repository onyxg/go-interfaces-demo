package domain

type Person interface {
	GetName() string
	GetAge() int
	NeedsToShower() bool
}

type SuperHero interface {
	Person
	GetHeroName() string
	CanFly() bool
	IsHuman() bool
	IsSuperSpeed() bool
	IsSuperStrength() bool
}
