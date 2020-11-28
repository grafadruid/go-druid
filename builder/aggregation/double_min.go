package aggregation

type DoubleMin struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewDoubleMin() *DoubleMin {
	d := &DoubleMin{}
	d.SetType("doubleMin")
	return d
}

func (d *DoubleMin) SetName(name string) *DoubleMin {
	d.Base.SetName(name)
	return d
}

func (d *DoubleMin) SetFieldName(fieldName string) *DoubleMin {
	d.FieldName = fieldName
	return d
}

func (d *DoubleMin) SetExpression(expression string) *DoubleMin {
	d.Expression = expression
	return d
}
