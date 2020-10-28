package filter

type True struct {
	Base
}

func NewTrue() *True {
	t := &True{}
	t.SetType("true")
	return t
}
