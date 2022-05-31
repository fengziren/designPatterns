package realization

type Rectangle struct{}

func (*Rectangle) Draw() {
	println("draw rectangle")
}
