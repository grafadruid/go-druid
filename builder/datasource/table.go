package datasource

type Table struct {
	Base
	Name string `json:"name,omitempty"`
}

func NewTable() *Table {
	t := &Table{}
	t.SetType("table")
	return t
}

func (t *Table) SetName(name string) *Table {
	t.Name = name
	return t
}
