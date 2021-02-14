package granularity

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
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

func Load(data []byte) (builder.Granularity, error) {
	var g builder.Granularity
	var t struct {
		Typ string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		g = NewSimple()
		return g, json.Unmarshal(data, &g)
	}
	switch t.Typ {
	case "duration":
		g = NewDuration()
	case "period":
		g = NewPeriod()
	case "all", "none":
		g = NewComplexSimple()
	default:
		return nil, errors.New("unsupported type")
	}
	return g, json.Unmarshal(data, &g)
}
