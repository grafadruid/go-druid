package intervals

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

func Load(data []byte) (builder.Intervals, error) {
	var i builder.Intervals
	if string(data) == "null" {
		return i, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "intervals":
		i = NewIntervals()
	default:
		return nil, errors.New("unsupported intervals type")
	}
	return i, json.Unmarshal(data, &i)
}
