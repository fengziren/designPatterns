package realization

import (
	"github.com/fengziren/designPatterns/builder/meal"
	"github.com/fengziren/designPatterns/builder/meal/item/burger"
	"github.com/fengziren/designPatterns/builder/meal/item/colddrink"
)

type MealBuilder struct{}

func (m *MealBuilder) PrepareVegMeal() *meal.Meal {
	meal := new(meal.Meal)
	meal.AddItem(new(burger.VegBurger))
	meal.AddItem(new(colddrink.Coke))
	return meal
}

func (*MealBuilder) PrepareNonVegMeal() *meal.Meal {
	meal := new(meal.Meal)
	meal.AddItem(new(burger.ChickenBurger))
	meal.AddItem(new(colddrink.Pepsi))
	return meal
}
