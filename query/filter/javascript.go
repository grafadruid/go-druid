package filter

import "github.com/grafadruid/go-druid/query"

type Javascript struct {
	Base
	Dimension    string             `json:"dimension"`
	Function     string             `json:"function"`
	ExtractionFn query.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning      `json:"filterTuning,omitempty"`
}

func NewJavascript() *Javascript {
	j := &Javascript{}
	j.SetType("javascript")
	return j
}

func (j *Javascript) SetDimension(dimension string) *Javascript {
	j.Dimension = dimension
	return j
}

func (j *Javascript) SetFunction(function string) *Javascript {
	j.Function = function
	return j
}

func (j *Javascript) SetExtractionFn(extractionFn query.ExtractionFn) *Javascript {
	j.ExtractionFn = extractionFn
	return j
}

func (j *Javascript) SetFilterTuning(filterTuning *FilterTuning) *Javascript {
	j.FilterTuning = filterTuning
	return j
}
