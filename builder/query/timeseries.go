package query

import (
	"encoding/json"
	"errors"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/aggregation"
	"github.com/h2oai/go-druid/builder/filter"
	"github.com/h2oai/go-druid/builder/granularity"
	"github.com/h2oai/go-druid/builder/postaggregation"
	"github.com/h2oai/go-druid/builder/virtualcolumn"
)

type Timeseries struct {
	Base
	Descending       *bool                    `json:"descending,omitempty"`
	VirtualColumns   []builder.VirtualColumn  `json:"virtualColumns,omitempty"`
	Filter           builder.Filter           `json:"filter,omitempty"`
	Granularity      builder.Granularity      `json:"granularity,omitempty"`
	Aggregations     []builder.Aggregator     `json:"aggregations,omitempty"`
	PostAggregations []builder.PostAggregator `json:"postAggregations,omitempty"`
	Limit            int64                    `json:"limit,omitempty"`
}

func NewTimeseries() *Timeseries {
	t := &Timeseries{}
	t.SetQueryType("timeseries")
	return t
}

func (t *Timeseries) SetDataSource(dataSource builder.DataSource) *Timeseries {
	t.Base.SetDataSource(dataSource)
	return t
}

func (t *Timeseries) SetIntervals(intervals builder.Intervals) *Timeseries {
	t.Base.SetIntervals(intervals)
	return t
}

func (t *Timeseries) SetContext(context map[string]interface{}) *Timeseries {
	t.Base.SetContext(context)
	return t
}

func (t *Timeseries) SetDescending(descending bool) *Timeseries {
	t.Descending = &descending
	return t
}

func (t *Timeseries) SetVirtualColumns(virtualColumns []builder.VirtualColumn) *Timeseries {
	t.VirtualColumns = virtualColumns
	return t
}

func (t *Timeseries) SetFilter(filter builder.Filter) *Timeseries {
	t.Filter = filter
	return t
}

func (t *Timeseries) SetGranularity(granularity builder.Granularity) *Timeseries {
	t.Granularity = granularity
	return t
}

func (t *Timeseries) SetAggregations(aggregations []builder.Aggregator) *Timeseries {
	t.Aggregations = aggregations
	return t
}

func (t *Timeseries) SetPostAggregations(postAggregations []builder.PostAggregator) *Timeseries {
	t.PostAggregations = postAggregations
	return t
}

func (t *Timeseries) SetLimit(limit int64) *Timeseries {
	t.Limit = limit
	return t
}

func (t *Timeseries) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Descending       *bool             `json:"descending,omitempty"`
		VirtualColumns   []json.RawMessage `json:"virtualColumns,omitempty"`
		Filter           json.RawMessage   `json:"filter,omitempty"`
		Granularity      json.RawMessage   `json:"granularity,omitempty"`
		Aggregations     []json.RawMessage `json:"aggregations,omitempty"`
		PostAggregations []json.RawMessage `json:"postAggregations,omitempty"`
		Limit            int64             `json:"limit,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var v builder.VirtualColumn
	vv := make([]builder.VirtualColumn, len(tmp.VirtualColumns))
	for i := range tmp.VirtualColumns {
		if v, err = virtualcolumn.Load(tmp.VirtualColumns[i]); err != nil {
			err = errors.New("2")
			return err
		}
		vv[i] = v
	}
	var f builder.Filter
	if tmp.Filter != nil {
		f, err = filter.Load(tmp.Filter)
		if err != nil {
			return err
		}
	}
	gr, err := granularity.Load(tmp.Granularity)
	if err != nil {
		return err
	}
	var a builder.Aggregator
	aa := make([]builder.Aggregator, len(tmp.Aggregations))
	for i := range tmp.Aggregations {
		if a, err = aggregation.Load(tmp.Aggregations[i]); err != nil {
			return err
		}
		aa[i] = a
	}
	var p builder.PostAggregator
	pp := make([]builder.PostAggregator, len(tmp.PostAggregations))
	for i := range tmp.PostAggregations {
		if p, err = postaggregation.Load(tmp.PostAggregations[i]); err != nil {
			return err
		}
		pp[i] = p
	}
	err = t.Base.UnmarshalJSON(data)
	t.Descending = tmp.Descending
	t.VirtualColumns = vv
	t.Filter = f
	t.Granularity = gr
	t.Aggregations = aa
	t.PostAggregations = pp
	t.Limit = tmp.Limit
	return err
}
