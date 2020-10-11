package filter

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Bound struct {
	Base
	Dimension    string                 `json:"dimension"`
	Lower        string                 `json:"lower,omitempty"`
	Upper        string                 `json:"upper,omitempty"`
	LowerStrict  bool                   `json:"lowerStrict,omitempty"`
	UpperStrict  bool                   `json:"upperStrict,omitempty"`
	ExtractionFn query.ExtractionFn     `json:"extractionFn,omitempty"`
	Ordering     types.StringComparator `json:"ordering,omitempty"`
}

func NewBound() *Bound {
	b := &Bound{}
	b.SetType("bound")
	return b
}

func (b *Bound) SetDimension(dimension string) *Bound {
	b.Dimension = dimension
	return b
}

func (b *Bound) SetLower(lower string) *Bound {
	b.Lower = lower
	return b
}

func (b *Bound) SetUpper(upper string) *Bound {
	b.Upper = upper
	return b
}

func (b *Bound) SetLowerStrict(lowerStrict bool) *Bound {
	b.LowerStrict = lowerStrict
	return b
}

func (b *Bound) SetUpperStrict(upperStrict bool) *Bound {
	b.UpperStrict = upperStrict
	return b
}

func (b *Bound) SetExtractionFn(extractionFn query.ExtractionFn) *Bound {
	b.ExtractionFn = extractionFn
	return b
}

func (b *Bound) SetOrdering(ordering types.StringComparator) *Bound {
	b.Ordering = ordering
	return b
}
