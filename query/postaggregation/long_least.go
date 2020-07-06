package postaggregation

import "github.com/grafadruid/go-druid/query"

type LongLeast struct {
	*Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewLongLeast() *LongLeast {
	l := &LongLeast{}
	l.SetType("longLeast")
	return l
}

func (l *LongLeast) SetName(name string) *LongLeast {
	l.Base.SetName(name)
	return l
}

func (l *LongLeast) SetFields(fields []query.PostAggregator) *LongLeast {
	l.Fields = fields
	return l
}
