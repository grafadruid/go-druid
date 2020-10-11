package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Timeseries struct {
	Base
	Descending       bool                   `json:"descending"`
	VirtualColumns   []query.VirtualColumn  `json:"virtualColumns"`
	Filter           query.Filter           `json:"filter"`
	Granularity      query.Granularity      `json:"granularity"`
	Aggregations     []query.Aggregator     `json:"aggregations"`
	PostAggregations []query.PostAggregator `json:"postAggregations"`
	Limit            int64                  `json:"limit"`
}

func NewTimeseries() *Timeseries {
	t := &Timeseries{}
	t.SetQueryType("timeseries")
	return t
}

func (t *Timeseries) SetDataSource(dataSource query.DataSource) *Timeseries {
	t.Base.SetDataSource(dataSource)
	return t
}

func (t *Timeseries) SetIntervals(intervals []types.Interval) *Timeseries {
	t.Base.SetIntervals(intervals)
	return t
}

func (t *Timeseries) SetContext(context map[string]interface{}) *Timeseries {
	t.Base.SetContext(context)
	return t
}

func (t *Timeseries) SetDescending(descending bool) *Timeseries {
	t.Descending = descending
	return t
}

func (t *Timeseries) SetVirtualColumns(virtualColumns []query.VirtualColumn) *Timeseries {
	t.VirtualColumns = virtualColumns
	return t
}

func (t *Timeseries) SetFilter(filter query.Filter) *Timeseries {
	t.Filter = filter
	return t
}

func (t *Timeseries) SetGranularity(granularity query.Granularity) *Timeseries {
	t.Granularity = granularity
	return t
}

func (t *Timeseries) SetAggregations(aggregations []query.Aggregator) *Timeseries {
	t.Aggregations = aggregations
	return t
}

func (t *Timeseries) SetPostAggregations(postAggregations []query.PostAggregator) *Timeseries {
	t.PostAggregations = postAggregations
	return t
}

func (t *Timeseries) SetLimit(limit int64) *Timeseries {
	t.Limit = limit
	return t
}
