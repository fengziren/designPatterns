package pack

type Bottle struct{}

func (*Bottle) Pack() string {
	return "Bottle"
}
