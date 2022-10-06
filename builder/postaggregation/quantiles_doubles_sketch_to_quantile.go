package postaggregation

// QuantilesDoublesSketchToQuantile struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html#quantile
type QuantilesDoublesSketchToQuantile struct {
	Base
	Field    *QuantilesDoublesSketchField `json:"field,omitempty"`
	Fraction *float64                     `json:"fraction,omitempty"`
}

// NewQuantilesDoublesSketchToQuantile new instance of QuantilesDoublesSketchToQuantile
func NewQuantilesDoublesSketchToQuantile() *QuantilesDoublesSketchToQuantile {
	q := &QuantilesDoublesSketchToQuantile{}
	q.SetType("quantilesDoublesSketchToQuantile")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketchToQuantile) SetName(name string) *QuantilesDoublesSketchToQuantile {
	q.Base.SetName(name)
	return q
}

// SetFraction set fraction
func (q *QuantilesDoublesSketchToQuantile) SetFraction(fraction float64) *QuantilesDoublesSketchToQuantile {
	q.Fraction = &fraction
	return q
}

// SetField set QuantilesDoublesSketchField
func (q *QuantilesDoublesSketchToQuantile) SetField(field *QuantilesDoublesSketchField) *QuantilesDoublesSketchToQuantile {
	q.Field = field
	return q
}

// NewQuantilesDoublesSketchToQuantileField new instance of QuantilesDoublesSketchField
// Deprecated: Use NewQuantilesDoublesSketchField instead.
// TODO: This function is a duplicate of "func NewQuantilesDoublesSketchField()" to keep backward compatible
func NewQuantilesDoublesSketchToQuantileField() *QuantilesDoublesSketchField {
	return &QuantilesDoublesSketchField{}
}
