package datasource

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

func Load(data []byte) (builder.DataSource, error) {
	var t struct {
		Typ builder.ComponentType `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d builder.DataSource
	switch t.Typ {
	case "globalTable":
		d = NewGlobalTable()
	case "inline":
		d = NewInline()
	case "join":
		d = NewJoin()
	case "lookup":
		d = NewLookup()
	case "query":
		d = NewQuery()
	case "table":
		d = NewTable()
	case "union":
		d = NewUnion()
	}
	return d, json.Unmarshal(data, &d)
}
