package bojack

import (
	"fmt"
	"time"
)

type Bojack struct{}

func NewBojack() *Bojack {
	return &Bojack{}
}

func (b *Bojack) CallDianeAt3AM() {
	fmt.Println("*Ringing* Hey Diane... you up? I just... I dunno, everything sucks.")
}

func (b *Bojack) RuinRelationships() {
	fmt.Println("Oops, I said something I shouldn't have. Guess that friendship is over.")
}

func (b *Bojack) SleepAllDay() {
	fmt.Println("It's 3 PM? Who cares. *goes back to sleep*")
}

func (b *Bojack) DrinkAndForget() {
	fmt.Println("Time to drink until I don't feel feelings anymore.")
	time.Sleep(2 * time.Second)
	fmt.Println("Wait... what was I talking about?")
}

func (b *Bojack) AbleToDriveCar(place string, speedLimit int) error {
	fmt.Printf("I'm promising you, I can drive you to %s with a speed limit of %d km/h\n", place, speedLimit)

	// Simulate car problem
	if place == "texus" {
		return fmt.Errorf("I'm sorry, the car has a problem")
	}
	return nil
}
