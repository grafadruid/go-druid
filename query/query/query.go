package query

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Base struct {
	ID         string                 `json:"ID"`
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

func (b *Base) SetID(ID string) *Base {
	b.ID = ID
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

func (b *Base) Language() query.QueryLanguage {
	return query.NativeLanguage
}

func Load(qry []byte) (query.Query, error) {
	var b Base
	if err := json.Unmarshal(qry, &b); err != nil {
		return nil, err
	}
	var q query.Query
	switch b.QueryType {
	case "scan":
		q = NewScan()
	case "sql":
		q = NewSQLQuery()
	}
	json.Unmarshal(qry, &q)
	return q, nil
}
