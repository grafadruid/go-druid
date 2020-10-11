package filter

type T struct {
	Base
}

func NewT() *T {
	t := &T{}
	t.SetType("t")
	return t
}
