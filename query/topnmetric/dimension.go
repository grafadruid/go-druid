package topnmetric

import "github.com/grafadruid/go-druid/query/types"

type Dimension struct {
	*Base
	PreviousStop string                 `json:"previousStop"`
	Ordering     types.StringComparator `json:"ordering"`
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
