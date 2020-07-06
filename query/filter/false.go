package filter

type F struct {
	*Base
}

func NewF() *F {
	f := &F{}
	f.SetType("f")
	return f
}
