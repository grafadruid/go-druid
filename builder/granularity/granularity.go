package granularity

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ string `json:"type"`
}

func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.Granularity, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var g builder.Granularity
	switch t.Typ {
	case "duration":
		g = NewDuration()
	case "period":
		g = NewPeriod()
	default:
		g = NewSimple()
	}
	return g, json.Unmarshal(data, &g)
}
