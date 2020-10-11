package postaggregation

type FinalizingField struct {
	Base
	FieldName string `json:"fieldName"`
}

func NewFinalizingField() *FinalizingField {
	f := &FinalizingField{}
	f.SetType("finalizingField")
	return f
}

func (f *FinalizingField) SetName(name string) *FinalizingField {
	f.Base.SetName(name)
	return f
}

func (f *FinalizingField) SetFieldName(fieldName string) *FinalizingField {
	f.FieldName = fieldName
	return f
}
