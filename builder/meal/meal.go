package meal

import "fmt"

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(item Item) {
	m.items = append(m.items, item)
}

func (m *Meal) GetCost() float64 {
	var cost float64

	for _, item := range m.items {
		cost += item.Price()
	}

	return cost
}

func (m *Meal) ShowItems() {
	for _, item := range m.items {
		fmt.Println("Item :", item.Name(), ", Packing :", item.Packing().Pack(), ", Price :", item.Price())
	}
}
