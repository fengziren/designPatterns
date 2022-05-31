package realization

import (
	"github.com/fengziren/designPatterns/abstractfacctory/abstract/factory"
	shaper "github.com/fengziren/designPatterns/abstractfacctory/abstract/factory/shape"
)

type ShapeFactory struct{}

func (*ShapeFactory) GetShape(shape string) factory.Shape {
	switch shape {
	case "CIRCLE":
		return new(shaper.Circle)
	case "RECTANGLE":
		return new(shaper.Rectangle)
	case "SQUARE":
		return new(shaper.Square)
	default:
		return nil
	}
}

func (*ShapeFactory) GetColor(color string) factory.Color {
	return nil
}
