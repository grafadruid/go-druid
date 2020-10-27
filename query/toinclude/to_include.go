package toinclude

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Typ string `json:"type"`
}

func NewBase() *Base {
	b := &Base{}
	b.SetType("base")
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() query.ComponentType {
	return b.Typ
}

func Load(data []byte) (query.ToInclude, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var ti query.ToInclude
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
