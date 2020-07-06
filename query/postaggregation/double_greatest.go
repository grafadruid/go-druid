package postaggregation

import "github.com/grafadruid/go-druid/query"

type DoubleGreatest struct {
	*Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewDoubleGreatest() *DoubleGreatest {
	d := &DoubleGreatest{}
	d.SetType("doubleGreatest")
	return d
}

func (d *DoubleGreatest) SetName(name string) *DoubleGreatest {
	d.Base.SetName(name)
	return d
}

func (d *DoubleGreatest) SetFields(fields []query.PostAggregator) *DoubleGreatest {
	d.Fields = fields
	return d
}
