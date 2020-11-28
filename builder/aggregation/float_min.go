package aggregation

type FloatMin struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewFloatMin() *FloatMin {
	f := &FloatMin{}
	f.SetType("floatMin")
	return f
}

func (f *FloatMin) SetName(name string) *FloatMin {
	f.Base.SetName(name)
	return f
}

func (f *FloatMin) SetFieldName(fieldName string) *FloatMin {
	f.FieldName = fieldName
	return f
}

func (f *FloatMin) SetExpression(expression string) *FloatMin {
	f.Expression = expression
	return f
}
