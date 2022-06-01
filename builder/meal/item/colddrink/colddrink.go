package colddrink

import (
	"github.com/fengziren/designPatterns/builder/meal"
	"github.com/fengziren/designPatterns/builder/meal/pack"
)

type ColdDrink struct{}

func (*ColdDrink) Packing() meal.Packing {
	return new(pack.Bottle)
}

// func (*ColdDrink) Price() float64
