package color

import "fmt"

type Blue struct{}

func (*Blue) Fill() {
	fmt.Println("Inside Blue::fill() method.")
}
