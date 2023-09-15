package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/extractionfn"
)

type Regex struct {
	Base
	Dimension    string               `json:"dimension,omitempty"`
	Pattern      string               `json:"pattern,omitempty"`
	ExtractionFn builder.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning        `json:"filterTuning,omitempty"`
}

func NewRegex() *Regex {
	r := &Regex{}
	r.SetType("regex")
	return r
}

func (r *Regex) SetDimension(dimension string) *Regex {
	r.Dimension = dimension
	return r
}

func (r *Regex) SetPattern(pattern string) *Regex {
	r.Pattern = pattern
	return r
}

func (r *Regex) SetExtractionFn(extractionFn builder.ExtractionFn) *Regex {
	r.ExtractionFn = extractionFn
	return r
}

func (r *Regex) SetFilterTuning(filterTuning *FilterTuning) *Regex {
	r.FilterTuning = filterTuning
	return r
}

func (r *Regex) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Dimension    string          `json:"dimension,omitempty"`
		Pattern      string          `json:"pattern,omitempty"`
		ExtractionFn json.RawMessage `json:"extractionFn,omitempty"`
		FilterTuning *FilterTuning   `json:"filterTuning,omitempty"`
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
	r.Base = tmp.Base
	r.Dimension = tmp.Dimension
	r.Pattern = tmp.Pattern
	r.ExtractionFn = e
	r.FilterTuning = tmp.FilterTuning
	return nil
}
