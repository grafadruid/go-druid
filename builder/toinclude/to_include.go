package toinclude

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.ToInclude, error) {
	var t struct {
		Typ builder.ComponentType `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var ti builder.ToInclude
	switch t.Typ {
	case "all":
		ti = NewAll()
	case "list":
		ti = NewList()
	case "none":
		ti = NewNone()
	}
	return ti, json.Unmarshal(data, &ti)
}
