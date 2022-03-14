package postaggregation

// QuantilesDoublesSketchToQuantiles struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html#quantiles
type QuantilesDoublesSketchToQuantiles struct {
	Base
	Field     *QuantilesDoublesSketchField `json:"field,omitempty"`
	Fractions []float64                    `json:"fractions,omitempty"`
}

// NewQuantilesDoublesSketchToQuantiles new instance of QuantilesDoublesSketchToHistogram
func NewQuantilesDoublesSketchToQuantiles() *QuantilesDoublesSketchToQuantiles {
	q := &QuantilesDoublesSketchToQuantiles{}
	q.SetType("quantilesDoublesSketchToQuantiles")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketchToQuantiles) SetName(name string) *QuantilesDoublesSketchToQuantiles {
	q.Base.SetName(name)
	return q
}

// SetFractions set fractions
func (q *QuantilesDoublesSketchToQuantiles) SetFractions(fractions []float64) *QuantilesDoublesSketchToQuantiles {
	q.Fractions = fractions
	return q
}

// SetField set QuantilesDoublesSketchField
func (q *QuantilesDoublesSketchToQuantiles) SetField(field *QuantilesDoublesSketchField) *QuantilesDoublesSketchToQuantiles {
	q.Field = field
	return q
}
