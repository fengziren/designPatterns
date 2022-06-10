package protyp

type Clone interface {
	Clone() Clone
	GetID() uint64
	SetID(id uint64)
	GetTyp() string
	Draw()
}
