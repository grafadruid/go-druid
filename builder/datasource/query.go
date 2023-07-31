package datasource

import (
	"encoding/json"

	"github.com/adjoeio/go-druid/builder"
)

type Query struct {
	Base
	Query builder.Query `json:"query,omitempty"`
}

func NewQuery() *Query {
	q := &Query{}
	q.SetType("query")
	return q
}

func (q *Query) SetQuery(qry builder.Query) *Query {
	q.Query = qry
	return q
}

func (q *Query) UnmarshalJSONWithQueryLoader(data []byte, loader func(data []byte) (builder.Query, error)) error {
	var tmp struct {
		Base
		Query json.RawMessage `json:"query,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	qry, err := loader(tmp.Query)
	if err != nil {
		return err
	}
	q.Base = tmp.Base
	q.Query = qry
	return nil
}
