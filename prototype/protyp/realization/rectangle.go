package realization

import "github.com/fengziren/designPatterns/prototype/protyp"

type Rectangle struct {
	*protyp.Shape
}



func (*Rectangle) Draw() {
	println("Inside Rectangle::draw() method.")
}