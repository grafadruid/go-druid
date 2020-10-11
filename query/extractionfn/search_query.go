package extractionfn

import "github.com/grafadruid/go-druid/query"

type SearchQuery struct {
	Base
	Query query.SearchQuerySpec `json:"query,omitempty"`
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
