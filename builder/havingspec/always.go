package havingspec

type Always struct {
	Base
}

func NewAlways() *Always {
	a := &Always{}
	a.SetType("always")
	return a
}
