package bound

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Typ string `json:"type"`
}

func NewBase() *Base {
	b := &Base{}
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() query.ComponentType {
	return b.Typ
}

func Load(data []byte) (query.Aggregator, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var b query.Bound
	switch t.Typ {
	case "polygon":
		b = NewPolygon()
	case "radius":
		b = NewRadius()
	case "rectangular":
		b = NewRectangular()
	}
	return b, json.Unmarshal(data, &b)
}
