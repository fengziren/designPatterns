package burger

import (
	"github.com/fengziren/designPatterns/builder/meal"
	"github.com/fengziren/designPatterns/builder/meal/pack"
)

type Burger struct{}

func (*Burger) Packing() meal.Packing {
	return new(pack.Wrapper)
}

// func (*Burger) Price() float64
