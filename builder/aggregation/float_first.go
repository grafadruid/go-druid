package aggregation

type FloatFirst struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	TimeColumn string `json:"timeColumn,omitempty"`
}

func NewFloatFirst() *FloatFirst {
	f := &FloatFirst{}
	f.SetType("floatFirst")
	return f
}

func (f *FloatFirst) SetName(name string) *FloatFirst {
	f.Base.SetName(name)
	return f
}

func (f *FloatFirst) SetFieldName(fieldName string) *FloatFirst {
	f.FieldName = fieldName
	return f
}

func (f *FloatFirst) SetTimeColumn(timeColumn string) *FloatFirst {
	f.TimeColumn = timeColumn
	return f
}
