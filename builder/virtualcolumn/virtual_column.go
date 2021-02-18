package virtualcolumn

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
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d builder.Dimension
	switch t.Typ {
	case "expression":
		d = NewExpression()
	default:
		return nil, errors.New("unsupported type")
	}
	return d, json.Unmarshal(data, &d)
}
