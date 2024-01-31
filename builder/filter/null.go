package filter

import (
	"encoding/json"
)

type Null struct {
	Base
	Column string `json:"column,omitempty"`
}

func NewNull() *Null {
	s := &Null{}
	s.SetType("null")
	return s
}

func (s *Null) SetColumn(column string) *Null {
	s.Column = column
	return s
}

func (s *Null) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Column string `json:"column,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	s.Base = tmp.Base
	s.Column = tmp.Column
	return nil
}
