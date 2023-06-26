package aggregation

import "github.com/grafadruid/go-druid/builder"

// Cardinality Each individual element of the "fields" list can be a String or DimensionSpec.
// A String dimension in the fields list is equivalent to a DefaultDimensionSpec (no transformations).
type Cardinality struct {
	Base
	Fields []builder.DimensionSpec `json:"fields,omitempty"`
	ByRow  *bool                   `json:"byRow,omitempty"`
	Round  *bool                   `json:"round,omitempty"`
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

func (c *Cardinality) SetFields(fields []builder.DimensionSpec) *Cardinality {
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
