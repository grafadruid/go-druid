package havingspec

type EqualTo struct {
	Base
	Aggregation string   `json:"aggregation,omitempty"`
	Value       *float64 `json:"value,omitempty"`
}

func NewEqualTo() *EqualTo {
	e := &EqualTo{}
	e.SetType("equalTo")
	return e
}

func (e *EqualTo) SetAggregation(aggregation string) *EqualTo {
	e.Aggregation = aggregation
	return e
}

func (e *EqualTo) SetValue(value float64) *EqualTo {
	e.Value = &value
	return e
}
