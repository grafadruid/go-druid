package filter

type Not struct {
	*Base
	Field string `json:"field"`
}

func NewNot() *Not {
	n := &Not{}
	n.SetType("not")
	return n
}

func (n *Not) SetField(field string) *Not {
	n.Field = field
	return n
}
