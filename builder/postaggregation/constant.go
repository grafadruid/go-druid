package postaggregation

type Constant struct {
	Base
	Value *float64 `json:"value,omitempty"`
}

func NewConstant() *Constant {
	c := &Constant{}
	c.SetType("constant")
	return c
}

func (c *Constant) SetName(name string) *Constant {
	c.Base.SetName(name)
	return c
}

func (c *Constant) SetValue(value float64) *Constant {
	c.Value = &value
	return c
}
