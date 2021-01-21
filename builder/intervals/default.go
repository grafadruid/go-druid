package intervals

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

func Load(data []byte) (builder.Intervals, error) {
	var i builder.Intervals
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	// Simple do not have Base so it does not have builder.ComponentType, so Simple will
	// return an error if we try to unmarshal. So on error we try to Unmarshal as Simple hoping it is Simple.
	if err := json.Unmarshal(data, &t); err != nil {
		i = NewSimple()
		return i, json.Unmarshal(data, &i)
	}

	switch t.Typ {
	case "intervals":
		i = NewIntervals()
	default:
		return nil, errors.New("unsupported type")
	}
	return i, json.Unmarshal(data, &i)
}
