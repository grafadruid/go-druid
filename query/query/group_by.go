package query

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/filter"
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

func (s *Scan) UnmarshalJSON(data []byte) error {
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
	//var vv []query.VirtualColumns
	//for j := range c.VirtualColumns {
	//vv = append(vv, virtualcolumns.Load(j))
	//}
	var aa []query.Aggregation
	for j = range c.Aggregations {

	}
	f, err := filter.Load(tmp.Filter)
	if err != nil {
		return err
	}
	s.Base = tmp.Base
	s.ResultFormat = tmp.ResultFormat
	s.BatchSize = tmp.BatchSize
	s.Limit = tmp.Limit
	s.Columns = tmp.Columns
	s.Legacy = tmp.Legacy
	s.Order = tmp.Order
	s.SetFilter(f)
	//s.SetVirtualColumns(vv)
	return nil
}
