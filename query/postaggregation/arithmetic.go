package postaggregation

import "github.com/grafadruid/go-druid/query"

type Arithmetic struct {
	*Base
	Fn       string                 `json:"fn"`
	Fields   []query.PostAggregator `json:"fields"`
	Ordering string                 `json:"ordering"`
}

func NewArithmetic() *Arithmetic {
	a := &Arithmetic{}
	a.SetType("arithmetic")
	return a
}

func (a *Arithmetic) SetName(name string) *Arithmetic {
	a.Base.SetName(name)
	return a
}

func (a *Arithmetic) SetFn(fn string) *Arithmetic {
	a.Fn = fn
	return a
}

func (a *Arithmetic) SetFields(fields []query.PostAggregator) *Arithmetic {
	a.Fields = fields
	return a
}

func (a *Arithmetic) SetOrdering(ordering string) *Arithmetic {
	a.Ordering = ordering
	return a
}
