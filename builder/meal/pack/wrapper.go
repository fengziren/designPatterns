package pack

type Wrapper struct{}

func (*Wrapper) Pack() string {
	return "Wrapper"
}
