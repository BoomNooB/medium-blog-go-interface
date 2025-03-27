package diane

import (
	"fmt"
)

type Diane struct{}

func NewDiane() *Diane {
	return &Diane{}
}

// Implements the required methods

func (d *Diane) NotCheating() {
	fmt.Println("[Diane] I believe in honesty and hard work.")
}

func (d *Diane) CatPerson() int {
	fmt.Println("[Diane] I have three rescue cats.")
	return 3
}

func (d *Diane) CanCook(kind string) string {
	fmt.Printf("[Diane] I made a healthy version of %s. Hope you like it!\n", kind)
	return kind
}

func (d *Diane) AlwaysAnswerWithYesOrNo(question string) bool {
	fmt.Printf("[Diane] That's a deep question: %s... but okay, I'll say 'yes'.\n", question)
	return true
}

func (d *Diane) AbleToDriveCar(place string, speedLimit int) error {
	fmt.Printf("[Diane] Driving responsibly to %s at %d km/h.\n", place, speedLimit)
	return nil
}

// Extra methods that aren't in the interface

func (d *Diane) WriteBook(title string) {
	fmt.Printf("[Diane] I'm writing a book called '%s'. It'll be insightful!\n", title)
}

func (d *Diane) AdvocateForChange(cause string) {
	fmt.Printf("[Diane] Fighting for %s! The world needs change.\n", cause)
}

// The tricky part: parameter names are shuffled, but the types match!
func (d *Diane) SpellMagic(name string, food string, age int, pet string) {
	fmt.Printf("[Diane] Alright, let's cast a spell with %s, a %d-year-old %s, and some %s.\n", name, age, pet, food)
}
