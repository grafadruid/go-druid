package aggregation

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder"
)

type Generic struct {
	Base
	Fields map[string]interface{}
}

func NewGeneric(typ string) *Generic {
	g := &Generic{Fields: make(map[string]interface{})}
	g.SetType(typ)
	return g
}

func (g *Generic) SetField(name string, value interface{}) *Generic {
	switch name {
	case "type":
		g.SetType(value.(builder.ComponentType))
	case "name":
		g.SetName(value.(string))
	default:
		g.Fields[name] = value
	}
	return g
}

func (g *Generic) SetFields(fields map[string]interface{}) *Generic {
	g.Fields = fields
	return g
}

func (g *Generic) MergeFields(fields map[string]interface{}) *Generic {
	for name, value := range fields {
		g.SetField(name, value)
	}
	return g
}

func (g *Generic) MarshalJSON() ([]byte, error) {
	// Reuse Fields map to avoid allocating a new one if it fits the capacity
	if g.Typ != "" {
		g.Fields["type"] = g.Typ
	}
	if g.Name != "" {
		g.Fields["name"] = g.Name
	}

	data, err := json.Marshal(&g.Fields)

	delete(g.Fields, "type")
	delete(g.Fields, "name")

	return data, err
}

func (g *Generic) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &g.Fields)

	if typ, present := g.Fields["type"]; present {
		g.Typ = typ.(builder.ComponentType)
	}
	if name, present := g.Fields["name"]; present {
		g.Name = name.(string)
	}

	delete(g.Fields, "type")
	delete(g.Fields, "name")

	return err
}
