package datasource

import (
	"github.com/grafadruid/go-druid/query"
)

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

//func (q *Query) UnmarshalJSON(data []byte) error {
//var tmp struct {
//Base
//Query json.RawMessage `json:"query"`
//}
//if err := json.Unmarshal(data, &tmp); err != nil {
//return err
//}
//qry, err := queries.Load(tmp.Query)
//if err != nil {
//return err
//}
//q.Base = tmp.Base
//q.Query = qry
//return nil
//}
