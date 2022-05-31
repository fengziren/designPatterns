package shape

import "fmt"

type Circle struct{}

func (*Circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}
