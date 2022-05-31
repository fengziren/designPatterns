package color

import "fmt"

type Red struct{}

func (*Red) Fill() {
	fmt.Println("Inside Red::fill() method.")
}
