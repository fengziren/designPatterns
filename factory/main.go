package main

import sfactory "github.com/fengziren/designPatterns/factory/shapefactory"

func main() {

	// 得到工厂
	shapefactory := new(sfactory.ShapeFactory)
	// 得到一个形状
	shape1 := shapefactory.GetShape(sfactory.CIRCLE)
	// 调用draw方法
	shape1.Draw()

	// 得到一个形状
	shape2 := shapefactory.GetShape(sfactory.RECTANGLE)
	// 调用draw方法
	shape2.Draw()

	// 得到一个形状
	shape3 := shapefactory.GetShape(sfactory.SQUARE)
	// 调用draw方法
	shape3.Draw()
}
