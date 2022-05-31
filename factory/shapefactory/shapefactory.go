package shapefactory

import (
	"github.com/fengziren/designPatterns/factory/realization"
	"github.com/fengziren/designPatterns/factory/shape"
)

type shapeInternal string

const (
	CIRCLE    shapeInternal = "circle"
	SQUARE    shapeInternal = "square"
	RECTANGLE shapeInternal = "rectangle"
)

type ShapeFactory struct{}

func (*ShapeFactory) GetShape(shapeType shapeInternal) shape.Shape {
	if shapeType == CIRCLE {
		return new(realization.Circle)
	} else if shapeType == SQUARE {
		return new(realization.Square)
	} else if shapeType == RECTANGLE {
		return new(realization.Rectangle)
	} else {
		return nil
	}
}
