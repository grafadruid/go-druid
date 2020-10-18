package limitspec

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Type string
}

func NewBase() *Base {
	b := &Base{}
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Type = typ
	return b
}

func Load(data []byte) (query.Dimension, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d query.Dimension
	switch t.Typ {
	case "default":
		d = NewDefault()
	}
	return d, json.Unmarshal(data, &d)
}
