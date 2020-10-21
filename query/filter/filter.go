package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Typ string `json:"type"`
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

func (b *Base) Type() query.ComponentType {
	return b.Typ
}

func Load(data []byte) (query.Filter, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var f query.Filter
	switch t.Typ {
	case "and":
		f = NewAnd()
	case "bound":
		f = NewBound()
	case "columnComparison":
		f = NewColumnComparison()
	case "expression":
		f = NewExpression()
	case "extraction":
		f = NewExtraction()
	case "false":
		f = NewFalse()
	case "filterTuning":
		f = NewFilterTuning()
	case "in":
		f = NewIn()
	case "interval":
		f = NewInterval()
	case "javascript":
		f = NewJavascript()
	case "like":
		f = NewLike()
	case "not":
		f = NewNot()
	case "or":
		f = NewOr()
	case "regex":
		f = NewRegex()
	case "search":
		f = NewSearch()
	case "selector":
		f = NewSelector()
	case "spatial":
		f = NewSpatial()
	case "true":
		f = NewTrue()
	}
	return f, json.Unmarshal(data, &f)
}
