package filter

type Or struct {
	*Base
	Fields []string `json:"fields"`
}

func NewOr() *Or {
	o := &Or{}
	o.SetType("or")
	return o
}

func (o *Or) SetFields(fields []string) *Or {
	o.Fields = fields
	return o
}
