package todd

import (
	"fmt"
)

type Todd struct{}

func NewTodd() *Todd {
	return &Todd{}
}

func (t *Todd) AbleToDriveCar(speedLimit int, place string) error {
	fmt.Println("[Todd] I think i have to get someone car to drive.")
	fmt.Printf("[Todd] Driving at %d km/h to %s... wait, do I even have a car?\n", speedLimit, place)
	if place == "Disneyland" {
		return fmt.Errorf("Oh no, I got distracted and forgot to drive.")
	}
	return nil
}

func (t *Todd) CanCook(kind string) string {
	fmt.Printf("[Todd] I made %s! ...Accidentally. But still, enjoy!\n", kind)
	return kind
}

func (t *Todd) NotCheating() {
	fmt.Println("[Todd] Cheating? Nah, too much effort.")
}

func (t *Todd) CatPerson() int {
	fmt.Println("[Todd] I have a cat. Just one. That's enough.")
	return 1
}

func (t *Todd) AlwaysAnswerWithYesOrNo(question string) bool {
	fmt.Println("[Todd] Yes, OK, sure, whatever you say.")
	return true // Todd just agrees with everything.
}
