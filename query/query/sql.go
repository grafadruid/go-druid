package query

import (
        "github.com/grafadruid/go-druid/query"
)

type SQLQuery struct {
	Base
	Query        string              `json:"query"`
	ResultFormat string              `json:"resultFormat"`
	Header       bool                `json:"header"`
	Context      map[string]string   `json:"context"`
	Parameters   []SQLQueryParameter `json:"parameters"`
}

type SQLQueryParameter struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewSQLQuery() *SQLQuery {
	s := &SQLQuery{}
	return s
}

func NewSQLQueryParameter() *SQLQueryParameter {
	p := &SQLQueryParameter{}
	return p
}

func (s *SQLQuery) SetQuery(query string) *SQLQuery {
	s.Query = query
	return s
}

func (s *SQLQuery) SetResultFormat(resultFormat string) *SQLQuery {
	s.ResultFormat = resultFormat
	return s
}

func (s *SQLQuery) SetHeader(header bool) *SQLQuery {
	s.Header = header
	return s
}

func (s *SQLQuery) SetContext(context map[string]string) *SQLQuery {
	s.Context = context
	return s
}

func (s *SQLQuery) SetParameters(parameters []SQLQueryParameter) *SQLQuery {
	s.Parameters = parameters
	return s
}

func (s *SQLQuery) Language() query.QueryLanguage {
	return query.SQLLanguage
}
