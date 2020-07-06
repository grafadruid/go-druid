package postaggregation

type Base struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func NewBase() *Base {
	b := &Base{}
	b.SetType("base")
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Type = typ
	return b
}

func (b *Base) SetName(name string) *Base {
	b.Name = name
	return b
}
