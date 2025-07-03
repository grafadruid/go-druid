package postaggregation

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/common"
)

// Base embeds the shared NamedBase to eliminate code duplication
type Base struct {
	common.NamedBase
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
	case "thetaSketchEstimate":
		p = NewThetaSketchEstimate()
	case "":
		return nil, errors.New("missing postaggregation type")
	default:
		p = builder.NewSpec(t.Typ)
	}
	return p, json.Unmarshal(data, &p)
}
