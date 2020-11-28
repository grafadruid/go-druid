package aggregation

type LongSum struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewLongSum() *LongSum {
	l := &LongSum{}
	l.SetType("longSum")
	return l
}

func (l *LongSum) SetName(name string) *LongSum {
	l.Base.SetName(name)
	return l
}

func (l *LongSum) SetFieldName(fieldName string) *LongSum {
	l.FieldName = fieldName
	return l
}

func (l *LongSum) SetExpression(expression string) *LongSum {
	l.Expression = expression
	return l
}
