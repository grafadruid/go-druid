package filter

import (
	"encoding/json"
	"errors"

	"github.com/h2oai/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type,omitempty"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.Filter, error) {
	var f builder.Filter
	if string(data) == "null" {
		return f, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
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
	default:
		return nil, errors.New("unsupported filter type")
	}
	return f, json.Unmarshal(data, &f)
}
