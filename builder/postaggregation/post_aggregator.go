package postaggregation

import (
	"encoding/json"
	"errors"

	"github.com/h2oai/go-druid/builder"
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
	var p builder.PostAggregator
	if string(data) == "null" {
		return p, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
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
	case "quantileFromTDigestSketch":
		p = NewQuantileFromTDigestSketch()
	case "quantilesFromTDigestSketch":
		p = NewQuantilesFromTDigestSketch()
	case "quantilesDoublesSketchToQuantile":
		p = NewQuantilesDoublesSketchToQuantile()
	case "quantilesDoublesSketchToQuantiles":
		p = NewQuantilesDoublesSketchToQuantiles()
	case "quantilesDoublesSketchToHistogram":
		p = NewQuantilesDoublesSketchToHistogram()
	case "quantilesDoublesSketchToRank":
		p = NewQuantilesDoublesSketchToRank()
	case "quantilesDoublesSketchToCDF":
		p = NewQuantilesDoublesSketchToCDF()
	case "quantilesDoublesSketchToString":
		p = NewQuantilesDoublesSketchToString()
	case "":
		return nil, errors.New("missing postaggregation type")
	default:
		p = builder.NewSpec(t.Typ)
	}
	return p, json.Unmarshal(data, &p)
}
