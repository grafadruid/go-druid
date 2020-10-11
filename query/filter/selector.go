package filter

import "github.com/grafadruid/go-druid/query"

type Selector struct {
	Base
	Dimension    string             `json:"dimension"`
	Value        string             `json:"value"`
	ExtractionFn query.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning      `json:"filterTuning,omitempty"`
}

func NewSelector() *Selector {
	s := &Selector{}
	s.SetType("selector")
	return s
}

func (s *Selector) SetDimension(dimension string) *Selector {
	s.Dimension = dimension
	return s
}

func (s *Selector) SetValue(value string) *Selector {
	s.Value = value
	return s
}

func (s *Selector) SetExtractionFn(extractionFn query.ExtractionFn) *Selector {
	s.ExtractionFn = extractionFn
	return s
}

func (s *Selector) SetFilterTuning(filterTuning *FilterTuning) *Selector {
	s.FilterTuning = filterTuning
	return s
}
