package postaggregation

type FinalizingFieldAccess struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewFinalizingFieldAccess() *FinalizingFieldAccess {
	f := &FinalizingFieldAccess{}
	f.SetType("finalizingFieldAccess")
	return f
}

func (f *FinalizingFieldAccess) SetName(name string) *FinalizingFieldAccess {
	f.Base.SetName(name)
	return f
}

func (f *FinalizingFieldAccess) SetFieldName(fieldName string) *FinalizingFieldAccess {
	f.FieldName = fieldName
	return f
}
