package builder

import "encoding/json"

type JSON struct {
	Typ    ComponentType
	Fields map[string]interface{}
}

func NewJSON(typ string) *JSON {
	g := &JSON{Fields: make(map[string]interface{})}
	g.SetType(typ)
	return g
}

// Methods for compatibility

func (g *JSON) Type() ComponentType {
	return g.Typ
}

func (g *JSON) SetType(typ string) *JSON {
	g.Typ = typ
	return g
}

func (g *JSON) SetName(name string) *JSON {
	g.SetField("name", name)
	return g
}

// Generic methods

func (g *JSON) SetField(name string, value interface{}) *JSON {
	switch name {
	case "type":
		g.SetType(value.(ComponentType))
	default:
		g.Fields[name] = value
	}
	return g
}

func (g *JSON) SetFields(fields map[string]interface{}) *JSON {
	if typ, present := fields["type"]; present {
		g.SetType(typ.(ComponentType))
		delete(fields, "type")
	}

	g.Fields = fields
	return g
}

func (g *JSON) MergeFields(fields map[string]interface{}) *JSON {
	for name, value := range fields {
		g.SetField(name, value)
	}
	return g
}

func (g *JSON) MarshalJSON() ([]byte, error) {
	// Reuse Fields map to avoid allocating a new one if it fits the capacity
	if g.Typ != "" {
		g.Fields["type"] = g.Typ
	}

	data, err := json.Marshal(&g.Fields)

	delete(g.Fields, "type")

	return data, err
}

func (g *JSON) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &g.Fields)

	if typ, present := g.Fields["type"]; present {
		g.Typ = typ.(ComponentType)
	}

	delete(g.Fields, "type")

	return err
}
