package limitspec

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType
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
	case "default":
		d = NewDefault()
	default:
		return nil, errors.New("unsupported limitspec type")
	}
	return d, json.Unmarshal(data, &d)
}
