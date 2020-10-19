package extractionfn

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/searchqueryspec"
)

type SearchQuery struct {
	Base
	Query query.SearchQuerySpec `json:"query"`
}

func NewSearchQuery() *SearchQuery {
	s := &SearchQuery{}
	s.SetType("searchQuery")
	return s
}

func (s *SearchQuery) SetQuery(query query.SearchQuerySpec) *SearchQuery {
	s.Query = query
	return s
}

func (s *SearchQuery) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Query json.RawMessage `json:"query"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	q, err := searchqueryspec.Load(tmp.Query)
	if err != nil {
		return err
	}
	s.Base = tmp.Base
	s.Query = q
	return nil
}
