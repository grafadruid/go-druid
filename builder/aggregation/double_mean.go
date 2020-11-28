package aggregation

type DoubleMean struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewDoubleMean() *DoubleMean {
	d := &DoubleMean{}
	d.SetType("doubleMean")
	return d
}

func (d *DoubleMean) SetName(name string) *DoubleMean {
	d.Base.SetName(name)
	return d
}

func (d *DoubleMean) SetFieldName(fieldName string) *DoubleMean {
	d.FieldName = fieldName
	return d
}
