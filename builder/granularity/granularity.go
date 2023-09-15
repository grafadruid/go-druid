package granularity

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/h2oai/go-druid/builder"
)

// Base is the base for granularity.
type Base struct {
	Typ string `json:"type,omitempty"`
}

// SetType sets type.
func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

// Type returns the type.
func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

// Load converts the druid native query to builder.Granularity
func Load(data []byte) (builder.Granularity, error) {
	var g builder.Granularity
	if string(data) == "null" {
		return g, nil
	}
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
		g = NewSimple()
		data = []byte(strconv.Quote(t.Typ))
	default:
		return nil, errors.New("unsupported granularity type")
	}
	return g, json.Unmarshal(data, &g)
}
