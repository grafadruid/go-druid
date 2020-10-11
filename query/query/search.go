package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type SearchSortSpec struct {
	Type types.StringComparator `json:"type"`
}

type Search struct {
	Base
	Filter           query.Filter          `json:"filter"`
	Granularity      query.Granularity     `json:"granularity"`
	Limit            int64                 `json:"limit"`
	SearchDimensions []query.Dimension     `json:"dimensions"`
	Query            query.SearchQuerySpec `json:"query"`
	Sort             *SearchSortSpec       `json:"sort"`
}

func NewSearchSearch() *Search {
	s := &Search{}
	s.SetQueryType("searchSortSpec Search")
	return s
}

func (s *Search) SetDataSource(dataSource query.DataSource) *Search {
	s.Base.SetDataSource(dataSource)
	return s
}

func (s *Search) SetIntervals(intervals []types.Interval) *Search {
	s.Base.SetIntervals(intervals)
	return s
}

func (s *Search) SetContext(context map[string]interface{}) *Search {
	s.Base.SetContext(context)
	return s
}

func (s *Search) SetFilter(filter query.Filter) *Search {
	s.Filter = filter
	return s
}

func (s *Search) SetGranularity(granularity query.Granularity) *Search {
	s.Granularity = granularity
	return s
}

func (s *Search) SetLimit(limit int64) *Search {
	s.Limit = limit
	return s
}

func (s *Search) SetSearchDimensions(searchDimensions []query.Dimension) *Search {
	s.SearchDimensions = searchDimensions
	return s
}

func (s *Search) SetQuery(query query.SearchQuerySpec) *Search {
	s.Query = query
	return s
}

func (s *Search) SetSort(sort *SearchSortSpec) *Search {
	s.Sort = sort
	return s
}
