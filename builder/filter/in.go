package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/extractionfn"
)

type In struct {
	Base
	Dimension    string               `json:"dimension"`
	Values       []string             `json:"values"`
	ExtractionFn builder.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning        `json:"filterTuning,omitempty"`
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

func (i *In) SetExtractionFn(extractionFn builder.ExtractionFn) *In {
	i.ExtractionFn = extractionFn
	return i
}

func (i *In) SetFilterTuning(filterTuning *FilterTuning) *In {
	i.FilterTuning = filterTuning
	return i
}

func (i *In) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension"`
		Values       []string        `json:"values"`
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
	i.Base = tmp.Base
	i.Dimension = tmp.Dimension
	i.Values = tmp.Values
	i.ExtractionFn = e
	i.FilterTuning = tmp.FilterTuning
	return nil
}
