package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/extractionfn"
)

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

func (j *Javascript) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension"`
		Function     string          `json:"function"`
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
	j.Base = tmp.Base
	j.Dimension = tmp.Dimension
	j.Function = tmp.Function
	j.ExtractionFn = e
	j.FilterTuning = tmp.FilterTuning
	return nil
}
