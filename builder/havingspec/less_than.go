package havingspec

type LessThan struct {
	Base
	Aggregation string   `json:"aggregation,omitempty"`
	Value       *float64 `json:"value,omitempty"`
}

func NewLessThan() *LessThan {
	l := &LessThan{}
	l.SetType("lessThan")
	return l
}

func (l *LessThan) SetAggregation(aggregation string) *LessThan {
	l.Aggregation = aggregation
	return l
}

func (l *LessThan) SetValue(value float64) *LessThan {
	l.Value = &value
	return l
}
