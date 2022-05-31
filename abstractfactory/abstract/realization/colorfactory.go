package realization

import (
	"github.com/fengziren/designPatterns/abstractfacctory/abstract/factory"
	colorr "github.com/fengziren/designPatterns/abstractfacctory/abstract/factory/color"
)

type ColorFactory struct{}

func (*ColorFactory) GetColor(color string) factory.Color {
	if color == "RED" {
		return new(colorr.Red)
	} else if color == "GREEN" {
		return new(colorr.Green)
	} else if color == "BLUE" {
		return new(colorr.Blue)
	}
	return nil
}

func (*ColorFactory) GetShape(shape string) factory.Shape {
	return nil
}
