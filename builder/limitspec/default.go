package limitspec

import (
	"github.com/h2oai/go-druid/builder/types"
)

type Direction string

const (
	Ascending  Direction = "ASCENDING"
	Descending           = "DESCENDING"
)

type OrderByColumnSpec struct {
	Dimension      string                 `json:"dimension,omitempty"`
	Direction      Direction              `json:"direction,omitempty"`
	DimensionOrder types.StringComparator `json:"dimensionOrder,omitempty"`
}

type Default struct {
	Base
	Columns []OrderByColumnSpec `json:"columns,omitempty"`
	Offset  int                 `json:"offset,omitempty"`
	Limit   int                 `json:"limit,omitempty"`
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
