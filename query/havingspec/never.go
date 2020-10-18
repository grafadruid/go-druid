package havingspec

type Never struct {
	Base
}

func NewNever() *Never {
	n := &Never{}
	n.SetType("never")
	return n
}
