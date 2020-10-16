package datasource

import "encoding/json"

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

func Load(data []byte) (query.Datasource, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d query.Datasource
	switch t.Typ {
	case "dataSource":
		d = NewDataSource()
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
