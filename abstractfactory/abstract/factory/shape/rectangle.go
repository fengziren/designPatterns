package shape

import "fmt"

type Rectangle struct{}

func (*Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}
