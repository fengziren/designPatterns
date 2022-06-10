package realization

import (
	"fmt"

	"github.com/fengziren/designPatterns/prototype/protyp"
)

type rectangle struct {
	id   uint64
	typ  string
	Name string
}

func GetRectangle() *rectangle {
	return &rectangle{typ: "rectangle"}
}

func (r *rectangle) Clone() protyp.Clone {
	clone := *r
	return &clone
}

func (r *rectangle) GetID() uint64 {
	return r.id
}

func (r *rectangle) SetID(id uint64) {
	r.id = id
}

func (r *rectangle) GetTyp() string {
	return r.typ
}

func (r *rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}
