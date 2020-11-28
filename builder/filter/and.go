package filter

type And struct {
	Base
	Fields []string `json:"fields,omitempty"`
}

func NewAnd() *And {
	a := &And{}
	a.SetType("and")
	return a
}

func (a *And) SetFields(fields []string) *And {
	a.Fields = fields
	return a
}
