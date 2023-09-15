package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/extractionfn"
)

type Selector struct {
	Base
	Dimension    string               `json:"dimension,omitempty"`
	Value        string               `json:"value,omitempty"`
	ExtractionFn builder.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning        `json:"filterTuning,omitempty"`
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

func (s *Selector) SetExtractionFn(extractionFn builder.ExtractionFn) *Selector {
	s.ExtractionFn = extractionFn
	return s
}

func (s *Selector) SetFilterTuning(filterTuning *FilterTuning) *Selector {
	s.FilterTuning = filterTuning
	return s
}

func (s *Selector) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Dimension    string          `json:"dimension,omitempty"`
		Value        string          `json:"value,omitempty"`
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
	s.Base = tmp.Base
	s.Dimension = tmp.Dimension
	s.Value = tmp.Value
	s.ExtractionFn = e
	s.FilterTuning = tmp.FilterTuning
	return nil
}
