package pc

import (
	"fmt"
)

type PC struct{}

func NewPC() *PC {
	return &PC{}
}

func (p *PC) NotCheating() {
	fmt.Println("[PC] I don't have time for that! I have a million things to do!")
}

func (p *PC) CatPerson() int {
	fmt.Println("[PC] I have so many cats. So. Many. Cats.")
	fmt.Printf("[PC] Let me count... 1, 2, 3, 4, 5...\n")
	return 5 // Let's say she has 5 cats
}

func (p *PC) CanCook(kind string) string {
	fmt.Printf("[PC] Sure, I can cook %s, but I will probably multitask and forget to eat it myself.\n", kind)
	return kind
}

func (p *PC) AlwaysAnswerWithYesOrNo(question string) bool {
	fmt.Printf("[PC] You asked: %s\n", question)
	if question == "Do you want to go out for dinner?" {
		return false
	}
	return true
}

func (p *PC) AbleToDriveCar(place string, speedLimit int) error {
	fmt.Printf("[PC] Driving to %s, but also making business calls at the same time.\n", place)
	if speedLimit < 60 {
		fmt.Println("[PC] Ugh, too slow! But fine, I will follow the rules.")
		return nil
	}
	if speedLimit > 100 {
		fmt.Println("[PC] I think we have a problem. I'm not sure if I can drive that fast.")
		return fmt.Errorf("Something is wrong with the car")
	}
	return nil
}
