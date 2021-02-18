package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ  builder.ComponentType `json:"type,omitempty"`
	Name string                `json:"name,omitempty"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) SetName(name string) *Base {
	b.Name = name
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.PostAggregator, error) {
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var p builder.PostAggregator
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
	case "quantilesFromTDigestSketch":
		p = NewQuantilesFromTDigestSketch()
		
	}
	return p, json.Unmarshal(data, &p)
}
