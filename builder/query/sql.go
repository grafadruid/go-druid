package query

import (
	"encoding/json"
)

type SQL struct {
	Base
	Query          string         `json:"query,omitempty"`
	ResultFormat   string         `json:"resultFormat,omitempty"`
	Header         *bool          `json:"header,omitempty"`
	TypesHeader    *bool          `json:"typesHeader,omitempty"`
	SQLTypesHeader *bool          `json:"sqlTypesHeader,omitempty"`
	Parameters     []SQLParameter `json:"parameters,omitempty"`
}

type SQLParameter struct {
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

func NewSQL() *SQL {
	s := &SQL{}
	s.Base.SetQueryType("sql")
	return s
}

func NewSQLParameter(Type string, Value interface{}) SQLParameter {
	p := SQLParameter{
		Type:  Type,
		Value: Value,
	}
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

func (s *SQL) SetTypesHeader(typesHeader bool) *SQL {
	s.TypesHeader = &typesHeader
	return s
}

func (s *SQL) SetSQLTypesHeader(sqlTypesHeader bool) *SQL {
	s.SQLTypesHeader = &sqlTypesHeader
	return s
}

func (s *SQL) SetParameters(parameters []SQLParameter) *SQL {
	s.Parameters = parameters
	return s
}

func (s *SQL) SetContext(context map[string]interface{}) *SQL {
	s.Base.SetContext(context)
	return s
}

func (s *SQL) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Query          string         `json:"query,omitempty"`
		ResultFormat   string         `json:"resultFormat,omitempty"`
		Header         *bool          `json:"header,omitempty"`
		TypesHeader    *bool          `json:"typesHeader,omitempty"`
		SQLTypesHeader *bool          `json:"sqlTypesHeader,omitempty"`
		Parameters     []SQLParameter `json:"parameters,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	err = s.Base.UnmarshalJSON(data)
	s.Query = tmp.Query
	s.ResultFormat = tmp.ResultFormat
	s.Header = tmp.Header
	s.TypesHeader = tmp.TypesHeader
	s.SQLTypesHeader = tmp.SQLTypesHeader
	s.Parameters = tmp.Parameters
	return err
}
