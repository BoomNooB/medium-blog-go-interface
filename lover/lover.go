package lover

type Ability interface {
	NotCheating()
	CatPerson() int
	CanCook(kind string)
	IsMyMsgRead(msg string) bool
	DriveCar(place string, speedLimit int) error
}

type lover struct {
	partner Ability
}

func MyLove(p Ability) *lover {
	return &lover{
		partner: p,
	}
}

func (l *lover) DoWhatLoverDo() {
	l.partner.NotCheating()
	l.partner.CatPerson()
	l.partner.CanCook("breakfast")
	l.partner.IsMyMsgRead("I love you")
	l.partner.DriveCar("supermarket", 60)
}
