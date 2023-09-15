package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/extractionfn"
	"github.com/h2oai/go-druid/builder/types"
)

type Bound struct {
	Base
	Dimension    string                 `json:"dimension,omitempty"`
	Lower        string                 `json:"lower,omitempty"`
	Upper        string                 `json:"upper,omitempty"`
	LowerStrict  *bool                  `json:"lowerStrict,omitempty"`
	UpperStrict  *bool                  `json:"upperStrict,omitempty"`
	ExtractionFn builder.ExtractionFn   `json:"extractionFn,omitempty"`
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
	b.LowerStrict = &lowerStrict
	return b
}

func (b *Bound) SetUpperStrict(upperStrict bool) *Bound {
	b.UpperStrict = &upperStrict
	return b
}

func (b *Bound) SetExtractionFn(extractionFn builder.ExtractionFn) *Bound {
	b.ExtractionFn = extractionFn
	return b
}

func (b *Bound) SetOrdering(ordering types.StringComparator) *Bound {
	b.Ordering = ordering
	return b
}

func (b *Bound) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Dimension    string                 `json:"dimension,omitempty"`
		Lower        string                 `json:"lower,omitempty"`
		Upper        string                 `json:"upper,omitempty"`
		LowerStrict  *bool                  `json:"lowerStrict,omitempty"`
		UpperStrict  *bool                  `json:"upperStrict,omitempty"`
		ExtractionFn json.RawMessage        `json:"extractionFn,omitempty"`
		Ordering     types.StringComparator `json:"ordering,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var e builder.ExtractionFn
	if tmp.ExtractionFn != nil {
		e, err = extractionfn.Load(tmp.ExtractionFn)
		if err != nil {
			return err
		}
	}
	b.Base = tmp.Base
	b.Dimension = tmp.Dimension
	b.Lower = tmp.Lower
	b.Upper = tmp.Upper
	b.LowerStrict = tmp.LowerStrict
	b.UpperStrict = tmp.UpperStrict
	b.ExtractionFn = e
	b.Ordering = tmp.Ordering
	return nil
}
