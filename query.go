package druid

import (
	"github.com/grafadruid/go-druid/query"
	base "github.com/grafadruid/go-druid/query/query"
)

const (
	NativeQueryEndpoint = "druid/v2"
	SQLQueryEndpoint    = "druid/v2/sql"
)

type QueryService struct {
	client *Client
}

func (q *QueryService) Execute(qry query.Query, result interface{}) (*Response, error) {
	var path string
	switch lang := qry.Language(); lang {
	case query.NativeLanguage:
		path = NativeQueryEndpoint
	case query.SQLLanguage:
		path = SQLQueryEndpoint
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

//func (q *QueryService) Cancel(query query.Query) () {}

//func (q *QueryService) Candidates(query query.Query, result interface{}) (*Response, error) {}

func (q *QueryService) Load(data []byte) (query.Query, error) {
	return base.Load(data)
}
