package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Base struct {
	QueryType  string                 `json:"queryType"`
	DataSource query.DataSource       `json:"dataSource"`
	Intervals  []types.Interval       `json:"intervals"`
	Context    map[string]interface{} `json:"context"`
}

func NewBase() *Base {
	b := &Base{}
	b.SetQueryType("base")
	return b
}

func (b *Base) SetQueryType(queryType string) *Base {
	b.QueryType = queryType
	return b
}

func (b *Base) SetDataSource(dataSource query.DataSource) *Base {
	b.DataSource = dataSource
	return b
}

func (b *Base) SetIntervals(intervals []types.Interval) *Base {
	b.Intervals = intervals
	return b
}

func (b *Base) SetContext(context map[string]interface{}) *Base {
	b.Context = context
	return b
}

func (b *Base) Language() QueryLanguage {
	return NativeLanguage
}
