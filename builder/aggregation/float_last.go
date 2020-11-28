package aggregation

type FloatLast struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewFloatLast() *FloatLast {
	f := &FloatLast{}
	f.SetType("floatLast")
	return f
}

func (f *FloatLast) SetName(name string) *FloatLast {
	f.Base.SetName(name)
	return f
}

func (f *FloatLast) SetFieldName(fieldName string) *FloatLast {
	f.FieldName = fieldName
	return f
}
