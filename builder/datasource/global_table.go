package datasource

type GlobalTable struct {
	Base
	Name string `json:"name,omitempty"`
}

func NewGlobalTable() *GlobalTable {
	g := &GlobalTable{}
	g.SetType("globalTable")
	return g
}

func (g *GlobalTable) SetName(name string) *GlobalTable {
	g.Name = name
	return g
}
