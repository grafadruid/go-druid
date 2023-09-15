package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/extractionfn"
	"github.com/h2oai/go-druid/builder/intervals"
)

type Interval struct {
	Base
	Dimension    string                `json:"dimension,omitempty"`
	Intervals    []*intervals.Interval `json:"intervals,omitempty"`
	ExtractionFn builder.ExtractionFn  `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning         `json:"filterTuning,omitempty"`
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

func (i *Interval) SetIntervals(intervals []*intervals.Interval) *Interval {
	i.Intervals = intervals
	return i
}

func (i *Interval) SetExtractionFn(extractionFn builder.ExtractionFn) *Interval {
	i.ExtractionFn = extractionFn
	return i
}

func (i *Interval) SetFilterTuning(filterTuning *FilterTuning) *Interval {
	i.FilterTuning = filterTuning
	return i
}

func (i *Interval) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Dimension    string          `json:"dimension,omitempty"`
		Intervals    json.RawMessage `json:"intervals,omitempty"`
		ExtractionFn json.RawMessage `json:"extractionFn,omitempty"`
		FilterTuning *FilterTuning   `json:"filterTuning,omitempty"`
	}
	if err = json.Unmarshal(data,
		&tmp); err != nil {
		return err
	}
	var e builder.ExtractionFn
	if tmp.ExtractionFn != nil {
		e, err = extractionfn.Load(tmp.ExtractionFn)
		if err != nil {
			return err
		}
	}
	var ii []*intervals.Interval
	if tmp.Intervals != nil {
		err = json.Unmarshal(tmp.Intervals,
			&ii)
		if err != nil {
			return err
		}
	}
	i.Base = tmp.Base
	i.Dimension = tmp.Dimension
	i.Intervals = ii
	i.ExtractionFn = e
	i.FilterTuning = tmp.FilterTuning
	return nil
}
