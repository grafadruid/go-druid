package aggregation

type FloatSum struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewFloatSum() *FloatSum {
	f := &FloatSum{}
	f.SetType("floatSum")
	return f
}

func (f *FloatSum) SetName(name string) *FloatSum {
	f.Base.SetName(name)
	return f
}

func (f *FloatSum) SetFieldName(fieldName string) *FloatSum {
	f.FieldName = fieldName
	return f
}

func (f *FloatSum) SetExpression(expression string) *FloatSum {
	f.Expression = expression
	return f
}
