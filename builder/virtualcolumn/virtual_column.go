package virtualcolumn

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

func Load(data []byte) (builder.VirtualColumn, error) {
	var v builder.VirtualColumn
	if string(data) == "null" {
		return v, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "expression":
		v = NewExpression()
	default:
		return nil, errors.New("unsupported virtualcolumn type")
	}
	return v, json.Unmarshal(data, &v)
}
