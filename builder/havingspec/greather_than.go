package havingspec

type GreaterThan struct {
	Base
	Aggregation string   `json:"aggregation,omitempty"`
	Value       *float64 `json:"value,omitempty"`
}

func NewGreaterThan() *GreaterThan {
	g := &GreaterThan{}
	g.SetType("greaterThan")
	return g
}

func (g *GreaterThan) SetAggregation(aggregation string) *GreaterThan {
	g.Aggregation = aggregation
	return g
}

func (g *GreaterThan) SetValue(value float64) *GreaterThan {
	g.Value = &value
	return g
}
