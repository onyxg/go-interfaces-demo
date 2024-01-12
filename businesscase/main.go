package main

import (
	"fmt"
	"go-interfaces-demo/businesscase/domain"
	"go-interfaces-demo/businesscase/factory"
)

func main() {

	percyDoc := factory.NewAbbrPerson("Percy", 42, true, false)
	natPerson := factory.NewNaturalPerson("John", 21)

	fmt.Printf("%s: needs to shower: %v\n", natPerson.GetName(), natPerson.NeedsToShower())
	fmt.Printf("%s: needs to shower: %v\n", percyDoc.GetName(), percyDoc.NeedsToShower())

	superHero, ok := percyDoc.(domain.SuperHero)
	if !ok {
		fmt.Printf("%s: is not a super hero\n", percyDoc.GetName())
		return
	} else {
		fmt.Printf("%s: can fly: %v", superHero.GetName(), superHero.CanFly())
	}

}
