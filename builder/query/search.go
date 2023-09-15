package query

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/dimension"
	"github.com/h2oai/go-druid/builder/filter"
	"github.com/h2oai/go-druid/builder/granularity"
	"github.com/h2oai/go-druid/builder/searchqueryspec"
	"github.com/h2oai/go-druid/builder/types"
)

type SearchSortSpec struct {
	Type types.StringComparator `json:"type,omitempty"`
}

type Search struct {
	Base
	Filter           builder.Filter          `json:"filter,omitempty"`
	Granularity      builder.Granularity     `json:"granularity,omitempty"`
	Limit            int64                   `json:"limit,omitempty"`
	SearchDimensions []builder.Dimension     `json:"searchDimensions,omitempty"`
	Query            builder.SearchQuerySpec `json:"query,omitempty"`
	Sort             *SearchSortSpec         `json:"sort,omitempty"`
}

func NewSearch() *Search {
	s := &Search{}
	s.SetQueryType("search")
	return s
}

func (s *Search) SetDataSource(dataSource builder.DataSource) *Search {
	s.Base.SetDataSource(dataSource)
	return s
}

func (s *Search) SetIntervals(intervals builder.Intervals) *Search {
	s.Base.SetIntervals(intervals)
	return s
}

func (s *Search) SetContext(context map[string]interface{}) *Search {
	s.Base.SetContext(context)
	return s
}

func (s *Search) SetFilter(filter builder.Filter) *Search {
	s.Filter = filter
	return s
}

func (s *Search) SetGranularity(granularity builder.Granularity) *Search {
	s.Granularity = granularity
	return s
}

func (s *Search) SetLimit(limit int64) *Search {
	s.Limit = limit
	return s
}

func (s *Search) SetSearchDimensions(searchDimensions []builder.Dimension) *Search {
	s.SearchDimensions = searchDimensions
	return s
}

func (s *Search) SetQuery(q builder.SearchQuerySpec) *Search {
	s.Query = q
	return s
}

func (s *Search) SetSort(sort *SearchSortSpec) *Search {
	s.Sort = sort
	return s
}

func (s *Search) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Filter           json.RawMessage   `json:"filter,omitempty"`
		Granularity      json.RawMessage   `json:"granularity,omitempty"`
		Limit            int64             `json:"limit,omitempty"`
		SearchDimensions []json.RawMessage `json:"searchDimensions,omitempty"`
		Query            json.RawMessage   `json:"query,omitempty"`
		Sort             *SearchSortSpec   `json:"sort,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var f builder.Filter
	if tmp.Filter != nil {
		f, err = filter.Load(tmp.Filter)
		if err != nil {
			return err
		}
	}
	var gr builder.Granularity
	if tmp.Granularity != nil {
		gr, err = granularity.Load(tmp.Granularity)
		if err != nil {
			return err
		}
	}
	var se builder.Dimension
	ss := make([]builder.Dimension, len(tmp.SearchDimensions))
	for i := range tmp.SearchDimensions {
		if se, err = dimension.Load(tmp.SearchDimensions[i]); err != nil {
			return err
		}
		ss[i] = se
	}
	q, err := searchqueryspec.Load(tmp.Query)
	if err != nil {
		return err
	}
	err = s.Base.UnmarshalJSON(data)
	s.Filter = f
	s.Granularity = gr
	s.SearchDimensions = ss
	s.Query = q
	s.Sort = tmp.Sort
	return err
}
