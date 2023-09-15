package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/bound"
)

type Spatial struct {
	Base
	Dimension    string        `json:"dimension,omitempty"`
	Bound        builder.Bound `json:"bound,omitempty"`
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

func (s *Spatial) SetBound(bound builder.Bound) *Spatial {
	s.Bound = bound
	return s
}

func (s *Spatial) SetFilterTuning(filterTuning *FilterTuning) *Spatial {
	s.FilterTuning = filterTuning
	return s
}

func (s *Spatial) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension,omitempty"`
		Bound        json.RawMessage `json:"bound,omitempty"`
		FilterTuning *FilterTuning   `json:"filterTuning,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	b, err := bound.Load(tmp.Bound)
	if err != nil {
		return err
	}
	s.Base = tmp.Base
	s.Dimension = tmp.Dimension
	s.Bound = b
	s.FilterTuning = tmp.FilterTuning
	return nil
}
