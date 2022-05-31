package abstract

import "github.com/fengziren/designPatterns/abstractfacctory/abstract/realization"

type FactoryProducer struct{}

func (f *FactoryProducer) GetFactory(factoryType string) AbstractFactory {
	switch factoryType {
	case "shape":
		return new(realization.ShapeFactory)
	case "color":
		return new(realization.ColorFactory)
	default:
		return nil
	}
}
