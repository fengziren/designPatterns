package burger

type VegBurger struct {
	*Burger
}

func (v *VegBurger) Price() float64 {
	return 25.0
}

func (v *VegBurger) Name() string {
	return "Veg Burger"
}
