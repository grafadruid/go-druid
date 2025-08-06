package aggregation

type LongFirst struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	TimeColumn string `json:"timeColumn,omitempty"`
}

func NewLongFirst() *LongFirst {
	l := &LongFirst{}
	l.SetType("longFirst")
	return l
}

func (l *LongFirst) SetName(name string) *LongFirst {
	l.Base.SetName(name)
	return l
}

func (l *LongFirst) SetFieldName(fieldName string) *LongFirst {
	l.FieldName = fieldName
	return l
}

func (l *LongFirst) SetTimeColumn(timeColumn string) *LongFirst {
	l.TimeColumn = timeColumn
	return l
}
