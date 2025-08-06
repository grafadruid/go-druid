package aggregation

type LongLast struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	TimeColumn string `json:"timeColumn,omitempty"`
}

func NewLongLast() *LongLast {
	l := &LongLast{}
	l.SetType("longLast")
	return l
}

func (l *LongLast) SetName(name string) *LongLast {
	l.Base.SetName(name)
	return l
}

func (l *LongLast) SetFieldName(fieldName string) *LongLast {
	l.FieldName = fieldName
	return l
}

func (l *LongLast) SetTimeColumn(timeColumn string) *LongLast {
	l.TimeColumn = timeColumn
	return l
}
