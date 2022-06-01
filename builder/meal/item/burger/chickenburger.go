package burger

type ChickenBurger struct {
	*Burger
}

func (*ChickenBurger) Price() float64 {
	return 50.50
}

func (*ChickenBurger) Name() string {
	return "Chicken Burger"
}
