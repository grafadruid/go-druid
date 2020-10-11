package postaggregation

import "github.com/grafadruid/go-druid/query"

type DoubleLeast struct {
	Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewDoubleLeast() *DoubleLeast {
	d := &DoubleLeast{}
	d.SetType("doubleLeast")
	return d
}

func (d *DoubleLeast) SetName(name string) *DoubleLeast {
	d.Base.SetName(name)
	return d
}

func (d *DoubleLeast) SetFields(fields []query.PostAggregator) *DoubleLeast {
	d.Fields = fields
	return d
}
