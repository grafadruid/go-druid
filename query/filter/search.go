package filter

import "github.com/grafadruid/go-druid/query"

type Search struct {
	*Base
	Dimension    string             `json:"dimension"`
	Query        string             `json:"query"`
	ExtractionFn query.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning      `json:"filterTuning,omitempty"`
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
