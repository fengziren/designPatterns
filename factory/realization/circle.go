package realization

type Circle struct{}

func (*Circle) Draw() {
	println("draw circle")
}
