package query

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/aggregation"
	"github.com/h2oai/go-druid/builder/dimension"
	"github.com/h2oai/go-druid/builder/filter"
	"github.com/h2oai/go-druid/builder/granularity"
	"github.com/h2oai/go-druid/builder/havingspec"
	"github.com/h2oai/go-druid/builder/limitspec"
	"github.com/h2oai/go-druid/builder/postaggregation"
	"github.com/h2oai/go-druid/builder/virtualcolumn"
)

type GroupBy struct {
	Base
	Dimensions       []builder.Dimension      `json:"dimensions,omitempty"`
	VirtualColumns   []builder.VirtualColumn  `json:"virtualColumns,omitempty"`
	Filter           builder.Filter           `json:"filter,omitempty"`
	Granularity      builder.Granularity      `json:"granularity,omitempty"`
	Aggregations     []builder.Aggregator     `json:"aggregations,omitempty"`
	PostAggregations []builder.PostAggregator `json:"postAggregations,omitempty"`
	Having           builder.HavingSpec       `json:"having,omitempty"`
	LimitSpec        builder.LimitSpec        `json:"limitSpec,omitempty"`
	SubtotalsSpec    [][]string               `json:"subtotalsSpec,omitempty"`
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

func (g *GroupBy) SetIntervals(intervals builder.Intervals) *GroupBy {
	g.Base.SetIntervals(intervals)
	return g
}

func (g *GroupBy) SetContext(context map[string]interface{}) *GroupBy {
	g.Base.SetContext(context)
	return g
}

func (g *GroupBy) SetDimensions(dimensions []builder.Dimension) *GroupBy {
	g.Dimensions = dimensions
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
	var err error
	var tmp struct {
		Dimensions       []json.RawMessage `json:"dimensions,omitempty"`
		VirtualColumns   []json.RawMessage `json:"virtualColumns,omitempty"`
		Filter           json.RawMessage   `json:"filter,omitempty"`
		Granularity      json.RawMessage   `json:"granularity,omitempty"`
		Aggregations     []json.RawMessage `json:"aggregations,omitempty"`
		PostAggregations []json.RawMessage `json:"postAggregations,omitempty"`
		Having           json.RawMessage   `json:"having,omitempty"`
		LimitSpec        json.RawMessage   `json:"limitSpec,omitempty"`
		SubtotalsSpec    [][]string        `json:"subtotalsSpec,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var d builder.Dimension
	dd := make([]builder.Dimension, len(tmp.Dimensions))
	for i := range tmp.Dimensions {
		if d, err = dimension.Load(tmp.Dimensions[i]); err != nil {
			return err
		}
		dd[i] = d
	}
	var v builder.VirtualColumn
	vv := make([]builder.VirtualColumn, len(tmp.VirtualColumns))
	for i := range tmp.VirtualColumns {
		if v, err = virtualcolumn.Load(tmp.VirtualColumns[i]); err != nil {
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
	var h builder.HavingSpec
	if tmp.Having != nil {
		h, err = havingspec.Load(tmp.Having)
		if err != nil {
			return err
		}
	}
	var l builder.LimitSpec
	if tmp.LimitSpec != nil {
		l, err = limitspec.Load(tmp.LimitSpec)
		if err != nil {
			return err
		}
	}
	if len(tmp.SubtotalsSpec) == 0 {
		tmp.SubtotalsSpec = nil
	}
	err = g.Base.UnmarshalJSON(data)
	g.Dimensions = dd
	g.VirtualColumns = vv
	g.Filter = f
	g.Granularity = gr
	g.Aggregations = aa
	g.PostAggregations = pp
	g.Having = h
	g.LimitSpec = l
	g.SubtotalsSpec = tmp.SubtotalsSpec
	return err
}
