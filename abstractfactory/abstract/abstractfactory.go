package abstract

import "github.com/fengziren/designPatterns/abstractfacctory/abstract/factory"

type AbstractFactory interface {
	GetColor(color string) factory.Color
	GetShape(shape string) factory.Shape
}
