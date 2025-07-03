package aggregation

type DoubleLast struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	TimeColumn string `json:"timeColumn,omitempty"`
}

func NewDoubleLast() *DoubleLast {
	d := &DoubleLast{}
	d.SetType("doubleLast")
	return d
}

func (d *DoubleLast) SetName(name string) *DoubleLast {
	d.Base.SetName(name)
	return d
}

func (d *DoubleLast) SetFieldName(fieldName string) *DoubleLast {
	d.FieldName = fieldName
	return d
}

func (l *DoubleLast) SetTimeColumn(timeColumn string) *DoubleLast {
	l.TimeColumn = timeColumn
	return l
}
