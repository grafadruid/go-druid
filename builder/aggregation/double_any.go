package aggregation

type DoubleAny struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewDoubleAny() *DoubleAny {
	d := &DoubleAny{}
	d.SetType("doubleAny")
	return d
}

func (d *DoubleAny) SetName(name string) *DoubleAny {
	d.Base.SetName(name)
	return d
}

func (d *DoubleAny) SetFieldName(fieldName string) *DoubleAny {
	d.FieldName = fieldName
	return d
}
