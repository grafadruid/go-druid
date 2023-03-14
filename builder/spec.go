package builder

import "encoding/json"

type Spec struct {
	Typ    ComponentType
	Fields map[string]interface{}
}

func NewSpec(typ string) *Spec {
	g := &Spec{Fields: make(map[string]interface{})}
	g.SetType(typ)
	return g
}

// Methods for compatibility

func (g *Spec) Type() ComponentType {
	return g.Typ
}

func (g *Spec) SetType(typ string) *Spec {
	g.Typ = typ
	return g
}

func (g *Spec) SetName(name string) *Spec {
	g.SetField("name", name)
	return g
}

// Generic methods

func (g *Spec) SetField(name string, value interface{}) *Spec {
	switch name {
	case "type":
		g.SetType(value.(ComponentType))
	default:
		g.Fields[name] = value
	}
	return g
}

func (g *Spec) SetFields(fields map[string]interface{}) *Spec {
	if typ, present := fields["type"]; present {
		g.SetType(typ.(ComponentType))
		delete(fields, "type")
	}

	g.Fields = fields
	return g
}

func (g *Spec) MergeFields(fields map[string]interface{}) *Spec {
	for name, value := range fields {
		g.SetField(name, value)
	}
	return g
}

func (g *Spec) MarshalJSON() ([]byte, error) {
	// Reuse Fields map to avoid allocating a new one if it fits the capacity
	if g.Typ != "" {
		g.Fields["type"] = g.Typ
	}

	data, err := json.Marshal(&g.Fields)

	delete(g.Fields, "type")

	return data, err
}

func (g *Spec) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &g.Fields)

	if typ, present := g.Fields["type"]; present {
		g.Typ = typ.(ComponentType)
	}

	delete(g.Fields, "type")

	return err
}
