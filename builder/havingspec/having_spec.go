package havingspec

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

func Load(data []byte) (builder.HavingSpec, error) {
	var h builder.HavingSpec
	if string(data) == "null" {
		return h, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "always":
		h = NewAlways()
	case "and":
		h = NewAnd()
	case "dimSelector":
		h = NewDimSelector()
	case "equalTo":
		h = NewEqualTo()
	case "greaterThan":
		h = NewGreaterThan()
	case "lessThan":
		h = NewLessThan()
	case "never":
		h = NewNever()
	case "not":
		h = NewNot()
	case "or":
		h = NewOr()
	default:
		return nil, errors.New("unsupported havingspec type")
	}
	return h, json.Unmarshal(data, &h)
}
