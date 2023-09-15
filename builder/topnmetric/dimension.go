package topnmetric

import "github.com/h2oai/go-druid/builder/types"

type Dimension struct {
	Base
	PreviousStop string                 `json:"previousStop,omitempty"`
	Ordering     types.StringComparator `json:"ordering,omitempty"`
}

func NewDimension() *Dimension {
	d := &Dimension{}
	d.SetType("dimension")
	return d
}

func (d *Dimension) SetPreviousStop(previousStop string) *Dimension {
	d.PreviousStop = previousStop
	return d
}

func (d *Dimension) SetOrdering(ordering types.StringComparator) *Dimension {
	d.Ordering = ordering
	return d
}
