package extractionfn

import (
	"github.com/grafadruid/go-druid/query/types"
)

type StringFormat struct {
	Base
	Format       string             `json:"format"`
	NullHandling types.NullHandling `json:"nullHandling"`
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
