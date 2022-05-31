package main

import "github.com/fengziren/designPatterns/abstractfacctory/abstract"

func main() {
	// 得到工厂生成器
	factoryProducer := abstract.FactoryProducer{}
	//获取形状工厂
	shapeFactory := factoryProducer.GetFactory("shape")
	// 获取形状
	circle := shapeFactory.GetShape("CIRCLE")
	// 调用draw方法
	circle.Draw()
	square := shapeFactory.GetShape("SQUARE")
	square.Draw()
	rectangle := shapeFactory.GetShape("RECTANGLE")
	rectangle.Draw()

	// 获取颜色工厂
	colorFactory := factoryProducer.GetFactory("color")
	// 获取颜色
	red := colorFactory.GetColor("RED")
	// 调用fill方法
	red.Fill()
	green := colorFactory.GetColor("GREEN")
	green.Fill()
	blue := colorFactory.GetColor("BLUE")
	blue.Fill()
}
