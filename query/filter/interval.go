package filter

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Interval struct {
	*Base
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
