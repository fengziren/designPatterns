package main

import (
	"fmt"

	"github.com/fengziren/designPatterns/builder/meal/realization"
)

func main() {
	mealBuilder := new(realization.MealBuilder)

	vegMeal := mealBuilder.PrepareVegMeal()
	fmt.Println("Veg Meal")
	vegMeal.ShowItems()
	fmt.Printf("Total Cost: %.2f\n", vegMeal.GetCost())

	nonVegMeal := mealBuilder.PrepareNonVegMeal()
	fmt.Println("\n\nNon-Veg Meal")
	nonVegMeal.ShowItems()
	fmt.Printf("Total Cost: %.2f\n", nonVegMeal.GetCost())
}
