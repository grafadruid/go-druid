package postaggregation

type FieldAccess struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewFieldAccess() *FieldAccess {
	f := &FieldAccess{}
	f.SetType("fieldAccess")
	return f
}

func (f *FieldAccess) SetName(name string) *FieldAccess {
	f.Base.SetName(name)
	return f
}

func (f *FieldAccess) SetFieldName(fieldName string) *FieldAccess {
	f.FieldName = fieldName
	return f
}
