package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type GroupBy struct {
	Base
	VirtualColumns   []query.VirtualColumn  `json:"virtualColumns"`
	Filter           query.Filter           `json:"filter"`
	Granularity      query.Granularity      `json:"granularity"`
	Aggregations     []query.Aggregator     `json:"aggregations"`
	PostAggregations []query.PostAggregator `json:"postAggregations"`
	Having           query.HavingSpec       `json:"having"`
	LimitSpec        query.LimitSpec        `json:"limitSpec"`
	SubtotalsSpec    [][]string             `json:"subtotalsSpec"`
}

func NewGroupBy() *GroupBy {
	g := &GroupBy{}
	g.SetQueryType("groupBy")
	return g
}

func (g *GroupBy) SetDataSource(dataSource query.DataSource) *GroupBy {
	g.Base.SetDataSource(dataSource)
	return g
}

func (g *GroupBy) SetIntervals(intervals []types.Interval) *GroupBy {
	g.Base.SetIntervals(intervals)
	return g
}

func (g *GroupBy) SetContext(context map[string]interface{}) *GroupBy {
	g.Base.SetContext(context)
	return g
}

func (g *GroupBy) SetVirtualColumns(virtualColumns []query.VirtualColumn) *GroupBy {
	g.VirtualColumns = virtualColumns
	return g
}

func (g *GroupBy) SetFilter(filter query.Filter) *GroupBy {
	g.Filter = filter
	return g
}

func (g *GroupBy) SetGranularity(granularity query.Granularity) *GroupBy {
	g.Granularity = granularity
	return g
}

func (g *GroupBy) SetAggregations(aggregations []query.Aggregator) *GroupBy {
	g.Aggregations = aggregations
	return g
}

func (g *GroupBy) SetPostAggregations(postAggregations []query.PostAggregator) *GroupBy {
	g.PostAggregations = postAggregations
	return g
}

func (g *GroupBy) SetHaving(having query.HavingSpec) *GroupBy {
	g.Having = having
	return g
}

func (g *GroupBy) SetLimitSpec(limitSpec query.LimitSpec) *GroupBy {
	g.LimitSpec = limitSpec
	return g
}

func (g *GroupBy) SetSubtotalsSpec(subtotalsSpec [][]string) *GroupBy {
	g.SubtotalsSpec = subtotalsSpec
	return g
}
