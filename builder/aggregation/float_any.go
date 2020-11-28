package aggregation

type FloatAny struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewFloatAny() *FloatAny {
	f := &FloatAny{}
	f.SetType("floatAny")
	return f
}

func (f *FloatAny) SetName(name string) *FloatAny {
	f.Base.SetName(name)
	return f
}

func (f *FloatAny) SetFieldName(fieldName string) *FloatAny {
	f.FieldName = fieldName
	return f
}
