package datasource

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

func Load(data []byte) (query.DataSource, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d query.DataSource
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
