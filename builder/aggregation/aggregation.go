package aggregation

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

func Load(data []byte) (builder.Aggregator, error) {
	var a builder.Aggregator
	if string(data) == "null" {
		return a, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "cardinality":
		a = NewCardinality()
	case "count":
		a = NewCount()
	case "doubleAny":
		a = NewDoubleAny()
	case "doubleFirst":
		a = NewDoubleFirst()
	case "doubleLast":
		a = NewDoubleLast()
	case "doubleMax":
		a = NewDoubleMax()
	case "doubleMean":
		a = NewDoubleMean()
	case "doubleMin":
		a = NewDoubleMin()
	case "doubleSum":
		a = NewDoubleSum()
	case "filtered":
		a = NewFiltered()
	case "floatAny":
		a = NewFloatAny()
	case "floatFirst":
		a = NewFloatFirst()
	case "floatLast":
		a = NewFloatLast()
	case "floatMax":
		a = NewFloatMax()
	case "floatMin":
		a = NewFloatMin()
	case "floatSum":
		a = NewFloatSum()
	case "histogram":
		a = NewHistogram()
	case "HLLSketchBuild":
		a = NewHLLSketchBuild()
	case "HLLSketchMerge":
		a = NewHLLSketchMerge()
	case "hyperUnique":
		a = NewHyperUnique()
	case "javascript":
		a = NewJavascript()
	case "longAny":
		a = NewLongAny()
	case "longFirst":
		a = NewLongFirst()
	case "longLast":
		a = NewLongLast()
	case "longMax":
		a = NewLongMax()
	case "longMin":
		a = NewLongMin()
	case "longSum":
		a = NewLongSum()
	case "stringAny":
		a = NewStringAny()
	case "stringFirstFolding":
		a = NewStringFirstFolding()
	case "stringFirst":
		a = NewStringFirst()
	case "stringLastFolding":
		a = NewStringLastFolding()
	case "stringLast":
		a = NewStringLast()
	case "tDigestSketch":
		a = NewTDigestSketch()
	case "quantilesDoublesSketch":
		a = NewQuantilesDoublesSketch()
	case "thetaSketch":
		a = NewThetaSketch()
	case "":
		return nil, errors.New("missing aggregation type")
	default:
		a = builder.NewSpec(t.Typ)
	}
	return a, json.Unmarshal(data, &a)
}
