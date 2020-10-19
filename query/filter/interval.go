package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/extractionfn"
	"github.com/grafadruid/go-druid/query/types"
)

type Interval struct {
	Base
	Dimension    string             `json:"dimension"`
	Intervals    []*types.Interval  `json:"intervals"`
	ExtractionFn query.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning      `json:"filterTuning,omitempty"`
}

func NewInterval() *Interval {
	i := &Interval{}
	i.SetType("interval")
	return i
}

func (i *Interval) SetDimension(dimension string) *Interval {
	i.Dimension = dimension
	return i
}

func (i *Interval) SetIntervals(intervals []*types.Interval) *Interval {
	i.Intervals = intervals
	return i
}

func (i *Interval) SetExtractionFn(extractionFn query.ExtractionFn) *Interval {
	i.ExtractionFn = extractionFn
	return i
}

func (i *Interval) SetFilterTuning(filterTuning *FilterTuning) *Interval {
	i.FilterTuning = filterTuning
	return i
}

func (i *Interval) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string            `json:"dimension"`
		Intervals    []*types.Interval `json:"intervals"`
		ExtractionFn json.RawMessage   `json:"extractionFn"`
		FilterTuning *FilterTuning     `json:"filterTuning"`
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
	i.Intervals = tmp.Intervals
	i.ExtractionFn = e
	i.FilterTuning = tmp.FilterTuning
	return nil
}
