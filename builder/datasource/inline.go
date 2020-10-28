package datasource

type Inline struct {
	Base
	ColumnNames []string   `json:"columnNames"`
	ColumnTypes []string   `json:"columnTypes"`
	Rows        [][]string `json:"rows"`
}

func NewInline() *Inline {
	i := &Inline{}
	i.SetType("inline")
	return i
}

func (i *Inline) SetColumnNames(columnNames []string) *Inline {
	i.ColumnNames = columnNames
	return i
}

func (i *Inline) SetColumnTypes(columnTypes []string) *Inline {
	i.ColumnTypes = columnTypes
	return i
}

func (i *Inline) SetRows(rows [][]string) *Inline {
	i.Rows = rows
	return i
}
