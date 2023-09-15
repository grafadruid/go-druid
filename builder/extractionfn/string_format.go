package extractionfn

import (
	"github.com/h2oai/go-druid/builder/types"
)

type StringFormat struct {
	Base
	Format       string             `json:"format,omitempty"`
	NullHandling types.NullHandling `json:"nullHandling,omitempty"`
}

func NewStringFormat() *StringFormat {
	s := &StringFormat{}
	s.SetType("stringFormat")
	return s
}

func (s *StringFormat) SetFormat(format string) *StringFormat {
	s.Format = format
	return s
}

func (s *StringFormat) SetNullHandling(nullHandling types.NullHandling) *StringFormat {
	s.NullHandling = nullHandling
	return s
}
