package aggregation

type LongMin struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewLongMin() *LongMin {
	l := &LongMin{}
	l.SetType("longMin")
	return l
}

func (l *LongMin) SetName(name string) *LongMin {
	l.Base.SetName(name)
	return l
}

func (l *LongMin) SetFieldName(fieldName string) *LongMin {
	l.FieldName = fieldName
	return l
}

func (l *LongMin) SetExpression(expression string) *LongMin {
	l.Expression = expression
	return l
}
