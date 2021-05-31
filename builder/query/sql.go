package query

import (
	"encoding/json"
)

type SQL struct {
	Base
	Query        string         `json:"query,omitempty"`
	ResultFormat string         `json:"resultFormat,omitempty"`
	Header       *bool          `json:"header,omitempty"`
	Parameters   []SQLParameter `json:"parameters,omitempty"`
}

type SQLParameter struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func NewSQL() *SQL {
	s := &SQL{}
	s.Base.SetQueryType("sql")
	return s
}

func NewSQLParameter() *SQLParameter {
	p := &SQLParameter{}
	return p
}

func (s *SQL) SetQuery(query string) *SQL {
	s.Query = query
	return s
}

func (s *SQL) SetResultFormat(resultFormat string) *SQL {
	s.ResultFormat = resultFormat
	return s
}

func (s *SQL) SetHeader(header bool) *SQL {
	s.Header = &header
	return s
}

func (s *SQL) SetParameters(parameters []SQLParameter) *SQL {
	s.Parameters = parameters
	return s
}

func (s *SQL) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Query        string         `json:"query,omitempty"`
		ResultFormat string         `json:"resultFormat,omitempty"`
		Header       *bool          `json:"header,omitempty"`
		Parameters   []SQLParameter `json:"parameters,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	err = s.Base.UnmarshalJSON(data)
	s.Query = tmp.Query
	s.ResultFormat = tmp.ResultFormat
	s.Header = tmp.Header
	s.Parameters = tmp.Parameters
	return err
}
