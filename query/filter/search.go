package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/extractionfn"
)

type Search struct {
	Base
	Dimension    string             `json:"dimension"`
	Query        string             `json:"query"`
	ExtractionFn query.ExtractionFn `json:"extractionFn"`
	FilterTuning *FilterTuning      `json:"filterTuning"`
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

func (s *Search) SetQuery(query string) *Search {
	s.Query = query
	return s
}

func (s *Search) SetExtractionFn(extractionFn query.ExtractionFn) *Search {
	s.ExtractionFn = extractionFn
	return s
}

func (s *Search) SetFilterTuning(filterTuning *FilterTuning) *Search {
	s.FilterTuning = filterTuning
	return s
}

func (s *Search) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension"`
		Query        string          `json:"query"`
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
	s.Base = tmp.Base
	s.Dimension = tmp.Dimension
	s.Query = tmp.Query
	s.ExtractionFn = e
	s.FilterTuning = tmp.FilterTuning
	return nil
}
