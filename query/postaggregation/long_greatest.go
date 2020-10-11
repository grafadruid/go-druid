package postaggregation

import "github.com/grafadruid/go-druid/query"

type LongGreatest struct {
	Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewLongGreatest() *LongGreatest {
	l := &LongGreatest{}
	l.SetType("longGreatest")
	return l
}

func (l *LongGreatest) SetName(name string) *LongGreatest {
	l.Base.SetName(name)
	return l
}

func (l *LongGreatest) SetFields(fields []query.PostAggregator) *LongGreatest {
	l.Fields = fields
	return l
}
