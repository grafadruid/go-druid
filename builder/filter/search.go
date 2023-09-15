package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/extractionfn"
)

type Search struct {
	Base
	Dimension    string               `json:"dimension,omitempty"`
	Query        string               `json:"builder,omitempty"`
	ExtractionFn builder.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning        `json:"filterTuning,omitempty"`
}

func NewSearch() *Search {
	s := &Search{}
	s.SetType("search")
	return s
}

func (s *Search) SetDimension(dimension string) *Search {
	s.Dimension = dimension
	return s
}

func (s *Search) SetQuery(q string) *Search {
	s.Query = q
	return s
}

func (s *Search) SetExtractionFn(extractionFn builder.ExtractionFn) *Search {
	s.ExtractionFn = extractionFn
	return s
}

func (s *Search) SetFilterTuning(filterTuning *FilterTuning) *Search {
	s.FilterTuning = filterTuning
	return s
}

func (s *Search) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Dimension    string          `json:"dimension,omitempty"`
		Query        string          `json:"builder,omitempty"`
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
	s.Query = tmp.Query
	s.ExtractionFn = e
	s.FilterTuning = tmp.FilterTuning
	return nil
}
