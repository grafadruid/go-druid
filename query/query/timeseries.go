package query

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/aggregation"
	"github.com/grafadruid/go-druid/query/filter"
	"github.com/grafadruid/go-druid/query/granularity"
	"github.com/grafadruid/go-druid/query/postaggregation"
	"github.com/grafadruid/go-druid/query/types"
	"github.com/grafadruid/go-druid/query/virtualcolumn"
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

func (t *Timeseries) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Descending      bool              `json:"descending"`
		VirtualColumns  []json.RawMessage `json:"virtualColumns"`
		Filter          json.RawMessage   `json:"filter"`
		Granularity     json.RawMessage   `json:"granularity"`
		Aggregators     []json.RawMessage `json:"aggregations"`
		PostAggregators []json.RawMessage `json:"postAggregations"`
		Limit           int64             `json:"limit"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var v query.VirtualColumn
	vv := make([]query.VirtualColumn, len(tmp.VirtualColumns))
	for i := range tmp.VirtualColumns {
		if v, err = virtualcolumn.Load(tmp.VirtualColumns[i]); err != nil {
			return err
		}
		vv[i] = v
	}
	f, err := filter.Load(tmp.Filter)
	if err != nil {
		return err
	}
	gr, err := granularity.Load(tmp.Granularity)
	if err != nil {
		return err
	}
	var a query.Aggregator
	aa := make([]query.Aggregator, len(tmp.Aggregators))
	for i := range tmp.Aggregators {
		if a, err = aggregation.Load(tmp.Aggregators[i]); err != nil {
			return err
		}
		aa[i] = a
	}
	var p query.PostAggregator
	pp := make([]query.PostAggregator, len(tmp.PostAggregators))
	for i := range tmp.PostAggregators {
		if p, err = postaggregation.Load(tmp.PostAggregators[i]); err != nil {
			return err
		}
		pp[i] = p
	}
	t.Base = tmp.Base
	t.Descending = tmp.Descending
	t.VirtualColumns = vv
	t.Filter = f
	t.Granularity = gr
	t.Aggregations = aa
	t.PostAggregations = pp
	t.Limit = tmp.Limit
	return nil
}
