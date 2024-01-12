package main

import (
	"go-interfaces-demo/businesscase/domain"
	"go-interfaces-demo/businesscase/factory"
	"log"
)

func main() {

	percyDoc := factory.NewAbbrPerson("Percy", 42, true, false)
	natPerson := factory.NewNaturalPerson("John", 21)

	log.Printf("%s: needs to shower: %v", natPerson.GetName(), natPerson.NeedsToShower())
	log.Printf("%s: needs to shower: %v", percyDoc.GetName(), percyDoc.NeedsToShower())

	// fmt.Println(percyDoc.GetName())
	// fmt.Println(natPerson.GetName())

	superHero, ok := percyDoc.(domain.SuperHero)
	if !ok {
		log.Printf("%s: is not a super hero", percyDoc.GetName())
		return
	} else {
		log.Printf("%s: can fly: %v", superHero.GetName(), superHero.CanFly())
	}

}
