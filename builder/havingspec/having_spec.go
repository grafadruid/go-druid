package havingspec

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
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

func Load(data []byte) (builder.Dimension, error) {
	var d builder.Dimension
	if string(data) == "null" {
		return d, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "always":
		d = NewAlways()
	case "and":
		d = NewAnd()
	case "dimSelector":
		d = NewDimSelector()
	case "equalTo":
		d = NewEqualTo()
	case "greaterThan":
		d = NewGreaterThan()
	case "lessThan":
		d = NewLessThan()
	case "never":
		d = NewNever()
	case "not":
		d = NewNot()
	case "or":
		d = NewOr()
	default:
		return nil, errors.New("unsupported havingspec type")
	}
	return d, json.Unmarshal(data, &d)
}
