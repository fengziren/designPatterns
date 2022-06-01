package colddrink

type Pepsi struct {
	*ColdDrink
}

func (*Pepsi) Price() float64 {
	return 35.00
}

func (*Pepsi) Name() string {
	return "Pepis"
}
