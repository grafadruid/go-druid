package intervals

import (
	"encoding/json"

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

func Load(data []byte) (builder.Intervals, error) {
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var i builder.Intervals
	switch t.Typ {
	case "intervals":
		i = NewIntervals()
	}
	return i, json.Unmarshal(data, &i)
}
