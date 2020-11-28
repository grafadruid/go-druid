package aggregation

type FloatMax struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewFloatMax() *FloatMax {
	f := &FloatMax{}
	f.SetType("floatMax")
	return f
}

func (f *FloatMax) SetName(name string) *FloatMax {
	f.Base.SetName(name)
	return f
}

func (f *FloatMax) SetFieldName(fieldName string) *FloatMax {
	f.FieldName = fieldName
	return f
}

func (f *FloatMax) SetExpression(expression string) *FloatMax {
	f.Expression = expression
	return f
}
