package aggregation

type DoubleSum struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewDoubleSum() *DoubleSum {
	d := &DoubleSum{}
	d.SetType("doubleSum")
	return d
}

func (d *DoubleSum) SetName(name string) *DoubleSum {
	d.Base.SetName(name)
	return d
}

func (d *DoubleSum) SetFieldName(fieldName string) *DoubleSum {
	d.FieldName = fieldName
	return d
}

func (d *DoubleSum) SetExpression(expression string) *DoubleSum {
	d.Expression = expression
	return d
}
