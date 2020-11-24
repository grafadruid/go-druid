package query

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/aggregation"
	"github.com/grafadruid/go-druid/builder/filter"
	"github.com/grafadruid/go-druid/builder/granularity"
	"github.com/grafadruid/go-druid/builder/havingspec"
	"github.com/grafadruid/go-druid/builder/limitspec"
	"github.com/grafadruid/go-druid/builder/postaggregation"
	"github.com/grafadruid/go-druid/builder/types"
	"github.com/grafadruid/go-druid/builder/virtualcolumn"
)

type GroupBy struct {
	Base
	VirtualColumns   []builder.VirtualColumn  `json:"virtualColumns"`
	Filter           builder.Filter           `json:"filter"`
	Granularity      builder.Granularity      `json:"granularity"`
	Aggregations     []builder.Aggregator     `json:"aggregations"`
	PostAggregations []builder.PostAggregator `json:"postAggregations"`
	Having           builder.HavingSpec       `json:"having"`
	LimitSpec        builder.LimitSpec        `json:"limitSpec"`
	SubtotalsSpec    [][]string               `json:"subtotalsSpec"`
}

func NewGroupBy() *GroupBy {
	g := &GroupBy{}
	g.SetQueryType("groupBy")
	return g
}

func (g *GroupBy) SetDataSource(dataSource builder.DataSource) *GroupBy {
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

func (g *GroupBy) SetVirtualColumns(virtualColumns []builder.VirtualColumn) *GroupBy {
	g.VirtualColumns = virtualColumns
	return g
}

func (g *GroupBy) SetFilter(filter builder.Filter) *GroupBy {
	g.Filter = filter
	return g
}

func (g *GroupBy) SetGranularity(granularity builder.Granularity) *GroupBy {
	g.Granularity = granularity
	return g
}

func (g *GroupBy) SetAggregations(aggregations []builder.Aggregator) *GroupBy {
	g.Aggregations = aggregations
	return g
}

func (g *GroupBy) SetPostAggregations(postAggregations []builder.PostAggregator) *GroupBy {
	g.PostAggregations = postAggregations
	return g
}

func (g *GroupBy) SetHaving(having builder.HavingSpec) *GroupBy {
	g.Having = having
	return g
}

func (g *GroupBy) SetLimitSpec(limitSpec builder.LimitSpec) *GroupBy {
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
		VirtualColumns   []json.RawMessage `json:"virtualColumns"`
		Filter           json.RawMessage   `json:"filter"`
		Granularity      json.RawMessage   `json:"granularity"`
		Aggregations     []json.RawMessage `json:"aggregations"`
		PostAggregations []json.RawMessage `json:"postAggregations"`
		Having           json.RawMessage   `json:"having"`
		LimitSpec        json.RawMessage   `json:"limitSpec"`
		SubtotalsSpec    [][]string        `json:"subtotalsSpec"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
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
