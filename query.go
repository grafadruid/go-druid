package druid

import (
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/query"
)

const (
	NativeQueryEndpoint = "druid/v2"
	SQLQueryEndpoint    = "druid/v2/sql"
)

type QueryService struct {
	client *Client
}

func (q *QueryService) Execute(qry builder.Query, result interface{}) (*Response, error) {
	var path string
	switch qry.Type() {
	case "sql":
		path = SQLQueryEndpoint
	default:
		path = NativeQueryEndpoint
	}
	r, err := q.client.NewRequest("POST", path, qry)
	if err != nil {
		return nil, err
	}
	resp, err := q.client.Do(r, result)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//func (q *QueryService) Cancel(query builder.Query) () {}

//func (q *QueryService) Candidates(query builder.Query, result interface{}) (*Response, error) {}

func (q *QueryService) Load(data []byte) (builder.Query, error) {
	return query.Load(data)
}
