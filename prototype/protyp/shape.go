package protyp

type Shape struct {
	id  string
	typ string
}

func (s *Shape) Draw()

func (s *Shape) GetType() string {
	return s.typ
}

func (s *Shape) GetId() string {
	return s.id
}

func (s *Shape) SetId(id string) {
	s.id = id
}

func (s *Shape) Clone() Clone {
	sCopy := *s
	return &sCopy
}

func (s *Shape) SetType(typ string) {
	s.typ = typ
}
