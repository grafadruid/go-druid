package searchqueryspec

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

func Load(data []byte) (query.SearchQuerySpec, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var s query.SearchQuerySpec
	switch t.Typ {
	case "all":
		s = NewAll()
	case "contains":
		s = NewContains()
	case "fragment":
		s = NewFragment()
	case "insensitiveContains":
		s = NewInsensitiveContains()
	case "regex":
		s = NewRegex()
	}
	return s, json.Unmarshal(data, &s)
}
