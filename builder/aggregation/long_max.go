package aggregation

type LongMax struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewLongMax() *LongMax {
	l := &LongMax{}
	l.SetType("longMax")
	return l
}

func (l *LongMax) SetName(name string) *LongMax {
	l.Base.SetName(name)
	return l
}

func (l *LongMax) SetFieldName(fieldName string) *LongMax {
	l.FieldName = fieldName
	return l
}

func (l *LongMax) SetExpression(expression string) *LongMax {
	l.Expression = expression
	return l
}
