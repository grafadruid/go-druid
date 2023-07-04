package postaggregation

// QuantilesDoublesSketchToString struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html#sketch-summary
type QuantilesDoublesSketchToString struct {
	Base
	Field *QuantilesDoublesSketchField `json:"field,omitempty"`
}

// NewQuantilesDoublesSketchToString new instance of QuantilesDoublesSketchToString
func NewQuantilesDoublesSketchToString() *QuantilesDoublesSketchToString {
	q := &QuantilesDoublesSketchToString{}
	q.SetType("quantilesDoublesSketchToString")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketchToString) SetName(name string) *QuantilesDoublesSketchToString {
	q.Base.SetName(name)
	return q
}

// SetField set QuantilesDoublesSketchField
func (q *QuantilesDoublesSketchToString) SetField(field *QuantilesDoublesSketchField) *QuantilesDoublesSketchToString {
	q.Field = field
	return q
}
