package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type TopN struct {
	*Base
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
