package main

import (
	"fmt"

	"github.com/fengziren/designPatterns/prototype/protyp/cache"
)

func main() {
	shapeCache := new(cache.ShapeCache)
	shapeCache.LoadCache()

	cloneShape := shapeCache.GetShape(1)
	fmt.Println("Shape: ", cloneShape.GetTyp())
	cloneShapec := cloneShape.Clone()
	fmt.Printf("%v : %p\t%v : %p\n", cloneShape, cloneShape, cloneShapec, cloneShapec)
}
