package extractionfn

type Strlen struct {
	Base
}

func NewStrlen() *Strlen {
	s := &Strlen{}
	s.SetType("strlen")
	return s
}
