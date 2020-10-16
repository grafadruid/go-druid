package granularity

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Type string `json:"type"`
}

func NewBase() *Base {
	b := &Base{}
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Type = typ
	return b
}

func Load(data []byte) (query.Granularity, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var g query.Granularity
	switch t.Typ {
	case "duration":
		g = NewDuration()
	case "period":
		g = NewPeriod()
	case "simple":
		g = NewSimple()
	}
	return g, json.Unmarshal(data, &g)
}
