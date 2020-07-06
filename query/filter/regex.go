package filter

import "github.com/grafadruid/go-druid/query"

type Regex struct {
	*Base
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
