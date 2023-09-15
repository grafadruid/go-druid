package query

import (
	"encoding/json"
	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/aggregation"
	"github.com/h2oai/go-druid/builder/dimension"
	"github.com/h2oai/go-druid/builder/filter"
	"github.com/h2oai/go-druid/builder/granularity"
	"github.com/h2oai/go-druid/builder/postaggregation"
	"github.com/h2oai/go-druid/builder/topnmetric"
	"github.com/h2oai/go-druid/builder/virtualcolumn"
)

type TopN struct {
	Base
	VirtualColumns   []builder.VirtualColumn  `json:"virtualColumns,omitempty"`
	Dimension        builder.Dimension        `json:"dimension,omitempty"`
	Metric           builder.TopNMetric       `json:"metric,omitempty"`
	Threshold        int64                    `json:"threshold,omitempty"`
	Filter           builder.Filter           `json:"filter,omitempty"`
	Granularity      builder.Granularity      `json:"granularity,omitempty"`
	Aggregations     []builder.Aggregator     `json:"aggregations,omitempty"`
	PostAggregations []builder.PostAggregator `json:"postAggregations,omitempty"`
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

func (t *TopN) SetIntervals(intervals builder.Intervals) *TopN {
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
	var err error
	var tmp struct {
		VirtualColumns   []json.RawMessage `json:"virtualColumns,omitempty"`
		Dimension        json.RawMessage   `json:"dimension,omitempty"`
		Metric           json.RawMessage   `json:"metric,omitempty"`
		Threshold        int64             `json:"threshold,omitempty"`
		Filter           json.RawMessage   `json:"filter,omitempty"`
		Granularity      json.RawMessage   `json:"granularity,omitempty"`
		Aggregations     []json.RawMessage `json:"aggregations,omitempty"`
		PostAggregations []json.RawMessage `json:"postAggregations,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
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
	d, err := dimension.Load(tmp.Dimension)
	if err != nil {
		return err
	}
	m, err := topnmetric.Load(tmp.Metric)
	if err != nil {
		return err
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
	t.VirtualColumns = vv
	t.Dimension = d
	t.Metric = m
	t.Threshold = tmp.Threshold
	t.Filter = f
	t.Granularity = gr
	t.Aggregations = aa
	t.PostAggregations = pp
	return err
}
