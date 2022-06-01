package colddrink

type Coke struct {
	*ColdDrink
}

func (*Coke) Price() float64 {
	return 30.00
}

func (*Coke) Name() string {
	return "Coke"
}
