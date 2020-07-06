package filter

import "github.com/grafadruid/go-druid/query"

type In struct {
	*Base
	Dimension    string             `json:"dimension"`
	Values       []string           `json:"values"`
	ExtractionFn query.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning      `json:"filterTuning,omitempty"`
}

func NewIn() *In {
	i := &In{}
	i.SetType("in")
	return i
}

func (i *In) SetDimension(dimension string) *In {
	i.Dimension = dimension
	return i
}

func (i *In) SetValues(values []string) *In {
	i.Values = values
	return i
}

func (i *In) SetExtractionFn(extractionFn query.ExtractionFn) *In {
	i.ExtractionFn = extractionFn
	return i
}

func (i *In) SetFilterTuning(filterTuning *FilterTuning) *In {
	i.FilterTuning = filterTuning
	return i
}
