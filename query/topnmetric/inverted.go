package topnmetric

import "github.com/grafadruid/go-druid/query"

type Inverted struct {
	Base
	Metric query.TopNMetric `json:"metric"`
}

func NewInverted() *Inverted {
	i := &Inverted{}
	i.SetType("inverted")
	return i
}

func (i *Inverted) SetMetric(metric query.TopNMetric) *Inverted {
	i.Metric = metric
	return i
}
