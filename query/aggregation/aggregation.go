package aggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func NewBase() *Base {
	b := &Base{}
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Type = typ
	return b
}

func (b *Base) SetName(name string) *Base {
	b.Name = name
	return b
}

func Load(data []byte) (query.Aggregator, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var a query.Aggregator
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
	}
	return a, json.Unmarshal(data, &a)
}
