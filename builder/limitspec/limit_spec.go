package limitspec

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

func Load(data []byte) (builder.LimitSpec, error) {
	var l builder.LimitSpec
	if string(data) == "null" {
		return l, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "default":
		l = NewDefault()
	default:
		return nil, errors.New("unsupported limitspec type")
	}
	return l, json.Unmarshal(data, &l)
}
