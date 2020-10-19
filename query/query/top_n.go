package query

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/aggregation"
	"github.com/grafadruid/go-druid/query/filter"
	"github.com/grafadruid/go-druid/query/granularity"
	"github.com/grafadruid/go-druid/query/postaggregation"
	"github.com/grafadruid/go-druid/query/topnmetric"
	"github.com/grafadruid/go-druid/query/types"
	"github.com/grafadruid/go-druid/query/virtualcolumn"
)

type TopN struct {
	Base
	VirtualColumns   []query.VirtualColumn  `json:"virtualColumns"`
	Dimension        query.Dimension        `json:"dimension"`
	Metric           query.TopNMetric       `json:"metric"`
	Threshold        int64                  `json:"threshold"`
	Filter           query.Filter           `json:"filter"`
	Granularity      query.Granularity      `json:"granularity"`
	Aggregations     []query.Aggregator     `json:"aggregations"`
	PostAggregations []query.PostAggregator `json:"postAggregations"`
}

func NewTopN() *TopN {
	t := &TopN{}
	t.SetQueryType("topN")
	return t
}

func (t *TopN) SetDataSource(dataSource query.DataSource) *TopN {
	t.Base.SetDataSource(dataSource)
	return t
}

func (t *TopN) SetIntervals(intervals []types.Interval) *TopN {
	t.Base.SetIntervals(intervals)
	return t
}

func (t *TopN) SetContext(context map[string]interface{}) *TopN {
	t.Base.SetContext(context)
	return t
}

func (t *TopN) SetVirtualColumns(virtualColumns []query.VirtualColumn) *TopN {
	t.VirtualColumns = virtualColumns
	return t
}

func (t *TopN) SetDimension(dimension query.Dimension) *TopN {
	t.Dimension = dimension
	return t
}

func (t *TopN) SetMetric(metric query.TopNMetric) *TopN {
	t.Metric = metric
	return t
}

func (t *TopN) SetThreshold(threshold int64) *TopN {
	t.Threshold = threshold
	return t
}

func (t *TopN) SetFilter(filter query.Filter) *TopN {
	t.Filter = filter
	return t
}

func (t *TopN) SetGranularity(granularity query.Granularity) *TopN {
	t.Granularity = granularity
	return t
}

func (t *TopN) SetAggregations(aggregations []query.Aggregator) *TopN {
	t.Aggregations = aggregations
	return t
}

func (t *TopN) SetPostAggregations(postAggregations []query.PostAggregator) *TopN {
	t.PostAggregations = postAggregations
	return t
}

func (t *TopN) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Metric          json.RawMessage   `json:"metric"`
		VirtualColumns  []json.RawMessage `json:"virtualColumns"`
		Threshold       int64             `json:"threshold"`
		Filter          json.RawMessage   `json:"filter"`
		Granularity     json.RawMessage   `json:"granularity"`
		Aggregators     []json.RawMessage `json:"aggregations"`
		PostAggregators []json.RawMessage `json:"postAggregations"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	m, err := topnmetric.Load(tmp.Metric)
	if err != nil {
		return err
	}
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
	t.Metric = m
	t.VirtualColumns = vv
	t.Threshold = tmp.Threshold
	t.Filter = f
	t.Granularity = gr
	t.Aggregations = aa
	t.PostAggregations = pp
	return nil
}
