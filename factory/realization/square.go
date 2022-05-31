package realization

type Square struct{}

func (*Square) Draw() {
	println("draw square")
}
