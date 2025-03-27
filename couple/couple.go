package couple

import "fmt"

type Ability interface {
	NotCheating()
	CatPerson() int
	CanCook(kind string) string
	AlwaysAnswerWithYesOrNo(question string) bool
	AbleToDriveCar(place string, speedLimit int) error
	SpellMagic(food string, pet string, age int, name string)
}

type couple struct {
	partner Ability
}

func NewCouple(partner Ability) *couple {
	return &couple{
		partner: partner,
	}
}

func (c *couple) DoWhatLoverDo() {
	fmt.Println("[ME] Will you cheat on me?")
	c.partner.NotCheating()

	fmt.Println("[ME] Babe, I'm so hungry, can you cook me Tom Yum Goong for me please 🥹 ?")

	food := c.partner.CanCook("Tom Yum Goong")
	fmt.Printf("[ME] Thanks, I'm sure that %s is delicious \n", food)

	q := "Do you miss me?"
	fmt.Printf("[ME] %s \n", q)
	answer := c.partner.AlwaysAnswerWithYesOrNo(q)
	if answer {
		fmt.Println("[ME] I miss you")
	} else {
		fmt.Println("[ME] Are you sure 🤬🤬🤬")
	}

	fmt.Println("[ME] Let's go to the supermarket, I want to buy some snacks")
	carHaveProblem := c.partner.AbleToDriveCar("supermarket", 60)
	if carHaveProblem != nil {
		fmt.Println("[ME] I think we should call a mechanic")
	} else {
		fmt.Println("[ME] Great, We finally arrived")
	}

	// SpellMagic(food string, pet string, age int, name string)
	fmt.Println("[ME] I need to be lucky person, can you cast a spell for me?")
	c.partner.SpellMagic("pizza", "cat", 2, "Luna")

	fmt.Println("[ME] I'm so happy to be with you, I love you so much ❤️❤️❤️")
}
