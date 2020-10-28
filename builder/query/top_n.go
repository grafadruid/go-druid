package query

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/aggregation"
	"github.com/grafadruid/go-druid/builder/filter"
	"github.com/grafadruid/go-druid/builder/granularity"
	"github.com/grafadruid/go-druid/builder/postaggregation"
	"github.com/grafadruid/go-druid/builder/topnmetric"
	"github.com/grafadruid/go-druid/builder/types"
	"github.com/grafadruid/go-druid/builder/virtualcolumn"
)

type TopN struct {
	Base
	VirtualColumns   []builder.VirtualColumn  `json:"virtualColumns"`
	Dimension        builder.Dimension        `json:"dimension"`
	Metric           builder.TopNMetric       `json:"metric"`
	Threshold        int64                    `json:"threshold"`
	Filter           builder.Filter           `json:"filter"`
	Granularity      builder.Granularity      `json:"granularity"`
	Aggregations     []builder.Aggregator     `json:"aggregations"`
	PostAggregations []builder.PostAggregator `json:"postAggregations"`
}

func NewTopN() *TopN {
	t := &TopN{}
	t.SetQueryType("topN")
	return t
}

func (t *TopN) SetDataSource(dataSource builder.DataSource) *TopN {
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

func (t *TopN) SetVirtualColumns(virtualColumns []builder.VirtualColumn) *TopN {
	t.VirtualColumns = virtualColumns
	return t
}

func (t *TopN) SetDimension(dimension builder.Dimension) *TopN {
	t.Dimension = dimension
	return t
}

func (t *TopN) SetMetric(metric builder.TopNMetric) *TopN {
	t.Metric = metric
	return t
}

func (t *TopN) SetThreshold(threshold int64) *TopN {
	t.Threshold = threshold
	return t
}

func (t *TopN) SetFilter(filter builder.Filter) *TopN {
	t.Filter = filter
	return t
}

func (t *TopN) SetGranularity(granularity builder.Granularity) *TopN {
	t.Granularity = granularity
	return t
}

func (t *TopN) SetAggregations(aggregations []builder.Aggregator) *TopN {
	t.Aggregations = aggregations
	return t
}

func (t *TopN) SetPostAggregations(postAggregations []builder.PostAggregator) *TopN {
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
	var v builder.VirtualColumn
	vv := make([]builder.VirtualColumn, len(tmp.VirtualColumns))
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
	var a builder.Aggregator
	aa := make([]builder.Aggregator, len(tmp.Aggregators))
	for i := range tmp.Aggregators {
		if a, err = aggregation.Load(tmp.Aggregators[i]); err != nil {
			return err
		}
		aa[i] = a
	}
	var p builder.PostAggregator
	pp := make([]builder.PostAggregator, len(tmp.PostAggregators))
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
