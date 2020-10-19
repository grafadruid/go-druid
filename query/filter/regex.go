package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/extractionfn"
)

type Regex struct {
	Base
	Dimension    string             `json:"dimension"`
	Pattern      string             `json:"pattern"`
	ExtractionFn query.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning      `json:"filterTuning,omitempty"`
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

func (r *Regex) SetExtractionFn(extractionFn query.ExtractionFn) *Regex {
	r.ExtractionFn = extractionFn
	return r
}

func (r *Regex) SetFilterTuning(filterTuning *FilterTuning) *Regex {
	r.FilterTuning = filterTuning
	return r
}

func (r *Regex) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension"`
		Pattern      string          `json:"pattern"`
		ExtractionFn json.RawMessage `json:"extractionFn"`
		FilterTuning *FilterTuning   `json:"filterTuning"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	e, err := extractionfn.Load(tmp.ExtractionFn)
	if err != nil {
		return err
	}
	r.Base = tmp.Base
	r.Dimension = tmp.Dimension
	r.Pattern = tmp.Pattern
	r.ExtractionFn = e
	r.FilterTuning = tmp.FilterTuning
	return nil
}
