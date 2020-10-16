package filter

type False struct {
	Base
}

func NewFalse() *False {
	f := &False{}
	f.SetType("false")
	return f
}
