package filter

import "github.com/grafadruid/go-druid/query"

type ColumnComparison struct {
	*Base
	Dimensions []query.Dimension `json:"dimensions"`
}

func NewColumnComparison() *ColumnComparison {
	c := &ColumnComparison{}
	c.SetType("columnComparison")
	return c
}

func (c *ColumnComparison) SetDimensions(dimensions []query.Dimension) *ColumnComparison {
	c.Dimensions = dimensions
	return c
}
