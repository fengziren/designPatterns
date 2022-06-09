package protyp

type Clone interface {
	GetType() string
	GetId() string
	SetId(id string)
	Clone() Clone
	Draw()
}
