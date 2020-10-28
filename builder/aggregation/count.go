package aggregation

type Count struct {
	Base
}

func NewCount() *Count {
	c := &Count{}
	c.SetType("count")
	return c
}

func (c *Count) SetName(name string) *Count {
	c.Base.SetName(name)
	return c
}
