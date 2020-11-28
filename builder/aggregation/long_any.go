package aggregation

type LongAny struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewLongAny() *LongAny {
	l := &LongAny{}
	l.SetType("longAny")
	return l
}

func (l *LongAny) SetName(name string) *LongAny {
	l.Base.SetName(name)
	return l
}

func (l *LongAny) SetFieldName(fieldName string) *LongAny {
	l.FieldName = fieldName
	return l
}
