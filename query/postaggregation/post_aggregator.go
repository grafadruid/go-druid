package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Typ  string `json:"type"`
	Name string `json:"name"`
}

func NewBase() *Base {
	b := &Base{}
	b.SetType("base")
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

func (b *Base) SetName(name string) *Base {
	b.Name = name
	return b
}

func (b *Base) Type() query.ComponentType {
	return b.Typ
}

func Load(data []byte) (query.PostAggregator, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var p query.PostAggregator
	switch t.Typ {
	case "arithmetic":
		p = NewArithmetic()
	case "constant":
		p = NewConstant()
	case "doubleGreatest":
		p = NewDoubleGreatest()
	case "doubleLeast":
		p = NewDoubleLeast()
	case "expression":
		p = NewExpression()
	case "fieldAccess":
		p = NewFieldAccess()
	case "finalizingFieldAccess":
		p = NewFinalizingFieldAccess()
	case "hyperUniqueFinalizing":
		p = NewHyperUniqueFinalizing()
	case "javascript":
		p = NewJavascript()
	case "longGreatest":
		p = NewLongGreatest()
	case "longLeast":
		p = NewLongLeast()
	}
	return p, json.Unmarshal(data, &p)
}
