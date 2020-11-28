package toinclude

type List struct {
	Base
	Columns []string `json:"columns,omitempty"`
}

func NewList() *List {
	l := &List{}
	l.SetType("list")
	return l
}

func (l *List) SetColumns(columns []string) *List {
	l.Columns = columns
	return l
}
