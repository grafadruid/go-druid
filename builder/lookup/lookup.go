package lookup

import (
	"encoding/json"
	"errors"

	"github.com/h2oai/go-druid/builder"
)

type Base struct {
	Typ string `json:"type,omitempty"`
}

func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.LookupExtractor, error) {
	var l builder.LookupExtractor
	if string(data) == "null" {
		return l, nil
	}
	var t struct {
		Typ string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "map":
		l = NewMap()
	default:
		return nil, errors.New("unsupported lookup type")
	}
	return l, json.Unmarshal(data, &l)
}
