package limitspec

import (
	"github.com/grafadruid/go-druid/builder/types"
)

type Direction string

const (
	Ascending  Direction = "ASCENDING"
	Descending           = "DESCENDING"
)

type OrderByColumnSpec struct {
	Dimension           string `json:"string"`
	Direction           Direction
	DimensionComparator types.StringComparator `json:"dimensionComparator"`
}

type Default struct {
	Base
	Columns []OrderByColumnSpec `json:"columns"`
	Offset  int                 `json:"offset"`
	Limit   int                 `json:"limit"`
}

func NewDefault() *Default {
	d := &Default{}
	d.SetType("default")
	return d
}

func (d *Default) SetColumns(columns []OrderByColumnSpec) *Default {
	d.Columns = columns
	return d
}

func (d *Default) SetOffset(offset int) *Default {
	d.Offset = offset
	return d
}

func (d *Default) SetLimit(limit int) *Default {
	d.Limit = limit
	return d
}
