package builder

import (
	"encoding/json"
)

type Generic struct {
	Typ    ComponentType
	Fields map[string]interface{}
}

func NewGeneric(typ string) *Generic {
	g := &Generic{Fields: make(map[string]interface{})}
	g.SetType(typ)
	return g
}

// Methods for compatibility

func (g *Generic) Type() ComponentType {
	return g.Typ
}

func (g *Generic) SetType(typ string) *Generic {
	g.Typ = typ
	return g
}

func (g *Generic) SetName(name string) *Generic {
	g.SetField("name", name)
	return g
}

// Generic methods

func (g *Generic) SetField(name string, value interface{}) *Generic {
	switch name {
	case "type":
		g.SetType(value.(ComponentType))
	default:
		g.Fields[name] = value
	}
	return g
}

func (g *Generic) SetFields(fields map[string]interface{}) *Generic {
	if typ, present := fields["type"]; present {
		g.SetType(typ.(ComponentType))
		delete(fields, "type")
	}

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

	data, err := json.Marshal(&g.Fields)

	delete(g.Fields, "type")

	return data, err
}

func (g *Generic) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &g.Fields)

	if typ, present := g.Fields["type"]; present {
		g.Typ = typ.(ComponentType)
	}

	delete(g.Fields, "type")

	return err
}
