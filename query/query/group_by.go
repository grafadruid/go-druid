package query

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/aggregation"
	"github.com/grafadruid/go-druid/query/filter"
	"github.com/grafadruid/go-druid/query/granularity"
	"github.com/grafadruid/go-druid/query/havingspec"
	"github.com/grafadruid/go-druid/query/limitspec"
	"github.com/grafadruid/go-druid/query/postaggregation"
	"github.com/grafadruid/go-druid/query/types"
	"github.com/grafadruid/go-druid/query/virtualcolumn"
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

func (g *GroupBy) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		VirtualColumns  []json.RawMessage `json:"virtualColumns"`
		Filter          json.RawMessage   `json:"filter"`
		Granularity     json.RawMessage   `json:"granularity"`
		Aggregators     []json.RawMessage `json:"aggregations"`
		PostAggregators []json.RawMessage `json:"postAggregations"`
		Having          json.RawMessage   `json:"having"`
		LimitSpec       json.RawMessage   `json:"limitSpec"`
		SubtotalsSpec   [][]string        `json:"subtotalsSpec"`
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
	h, err := havingspec.Load(tmp.Having)
	if err != nil {
		return err
	}
	l, err := limitspec.Load(tmp.LimitSpec)
	if err != nil {
		return err
	}
	g.Base = tmp.Base
	g.VirtualColumns = vv
	g.Filter = f
	g.Granularity = gr
	g.Aggregations = aa
	g.PostAggregations = pp
	g.Having = h
	g.LimitSpec = l
	g.SubtotalsSpec = tmp.SubtotalsSpec
	return nil
}
