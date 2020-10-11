package datasource

import "github.com/grafadruid/go-druid/query"

type Query struct {
	Base
	Query query.Query
}

func NewQuery() *Query {
	q := &Query{}
	q.SetType("query")
	return q
}

func (q *Query) SetQuery(query query.Query) {
	q.Query = query
}
