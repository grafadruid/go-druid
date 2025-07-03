package aggregation

type DoubleFirst struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	TimeColumn string `json:"timeColumn,omitempty"`
}

func NewDoubleFirst() *DoubleFirst {
	d := &DoubleFirst{}
	d.SetType("doubleFirst")
	return d
}

func (d *DoubleFirst) SetName(name string) *DoubleFirst {
	d.Base.SetName(name)
	return d
}

func (d *DoubleFirst) SetFieldName(fieldName string) *DoubleFirst {
	d.FieldName = fieldName
	return d
}

func (d *DoubleFirst) SetTimeColumn(timeColumn string) *DoubleFirst {
	d.TimeColumn = timeColumn
	return d
}
