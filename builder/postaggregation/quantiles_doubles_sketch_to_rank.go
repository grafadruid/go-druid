package postaggregation

// QuantilesDoublesSketchToRank struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html#rank
type QuantilesDoublesSketchToRank struct {
	Base
	Field *QuantilesDoublesSketchField `json:"field,omitempty"`
	Value *float64                     `json:"value,omitempty"`
}

// NewQuantilesDoublesSketchToRank new instance of NewQuantilesDoublesSketchToRank
func NewQuantilesDoublesSketchToRank() *QuantilesDoublesSketchToRank {
	q := &QuantilesDoublesSketchToRank{}
	q.SetType("quantilesDoublesSketchToRank")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketchToRank) SetName(name string) *QuantilesDoublesSketchToRank {
	q.Base.SetName(name)
	return q
}

// SetValue set value
func (q *QuantilesDoublesSketchToRank) SetValue(value float64) *QuantilesDoublesSketchToRank {
	q.Value = &value
	return q
}

// SetField set QuantilesDoublesSketchField
func (q *QuantilesDoublesSketchToRank) SetField(field *QuantilesDoublesSketchField) *QuantilesDoublesSketchToRank {
	q.Field = field
	return q
}
