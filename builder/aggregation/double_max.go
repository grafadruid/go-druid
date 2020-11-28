package aggregation

type DoubleMax struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func NewDoubleMax() *DoubleMax {
	d := &DoubleMax{}
	d.SetType("doubleMax")
	return d
}

func (d *DoubleMax) SetName(name string) *DoubleMax {
	d.Base.SetName(name)
	return d
}

func (d *DoubleMax) SetFieldName(fieldName string) *DoubleMax {
	d.FieldName = fieldName
	return d
}

func (d *DoubleMax) SetExpression(expression string) *DoubleMax {
	d.Expression = expression
	return d
}
