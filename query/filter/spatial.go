package filter

import "github.com/grafadruid/go-druid/query"

type Spatial struct {
	Base
	Dimension    string        `json:"dimension"`
	Bound        query.Bound   `json:"bound"`
	FilterTuning *FilterTuning `json:"filterTuning,omitempty"`
}

func NewSpatial() *Spatial {
	s := &Spatial{}
	s.SetType("spatial")
	return s
}

func (s *Spatial) SetDimension(dimension string) *Spatial {
	s.Dimension = dimension
	return s
}

func (s *Spatial) SetBound(bound query.Bound) *Spatial {
	s.Bound = bound
	return s
}

func (s *Spatial) SetFilterTuning(filterTuning *FilterTuning) *Spatial {
	s.FilterTuning = filterTuning
	return s
}
