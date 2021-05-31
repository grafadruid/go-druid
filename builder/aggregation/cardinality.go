package aggregation

type Cardinality struct {
	Base
	Fields []string `json:"fields,omitempty"`
	ByRow  *bool    `json:"byRow,omitempty"`
	Round  *bool    `json:"round,omitempty"`
}

func NewCardinality() *Cardinality {
	c := &Cardinality{}
	c.SetType("cardinality")
	return c
}

func (c *Cardinality) SetName(name string) *Cardinality {
	c.Base.SetName(name)
	return c
}

func (c *Cardinality) SetFields(fields []string) *Cardinality {
	c.Fields = fields
	return c
}

func (c *Cardinality) SetByRow(byRow bool) *Cardinality {
	c.ByRow = &byRow
	return c
}

func (c *Cardinality) SetRound(round bool) *Cardinality {
	c.Round = &round
	return c
}
