package query

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/filter"
)

type TimeBoundary struct {
	Base
	Bound  string         `json:"bound,omitempty"`
	Filter builder.Filter `json:"filter,omitempty"`
}

func NewTimeBoundary() *TimeBoundary {
	t := &TimeBoundary{}
	t.SetQueryType("timeBoundary")
	return t
}

func (t *TimeBoundary) SetDataSource(dataSource builder.DataSource) *TimeBoundary {
	t.Base.SetDataSource(dataSource)
	return t
}

func (t *TimeBoundary) SetIntervals(intervals builder.Intervals) *TimeBoundary {
	t.Base.SetIntervals(intervals)
	return t
}

func (t *TimeBoundary) SetContext(context map[string]interface{}) *TimeBoundary {
	t.Base.SetContext(context)
	return t
}

func (t *TimeBoundary) SetBound(bound string) *TimeBoundary {
	t.Bound = bound
	return t
}

func (t *TimeBoundary) SetFilter(filter builder.Filter) *TimeBoundary {
	t.Filter = filter
	return t
}

func (t *TimeBoundary) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Bound  string          `json:"bound,omitempty"`
		Filter json.RawMessage `json:"filter,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var f builder.Filter
	if tmp.Filter != nil {
		f, err = filter.Load(tmp.Filter)
		if err != nil {
			return err
		}
	}
	err = t.Base.UnmarshalJSON(data)
	t.Bound = tmp.Bound
	t.Filter = f
	return err
}
