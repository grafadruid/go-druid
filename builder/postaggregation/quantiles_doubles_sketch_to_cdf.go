package postaggregation

// QuantilesDoublesSketchToCDF struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html#cdf
type QuantilesDoublesSketchToCDF struct {
	Base
	Field       *QuantilesDoublesSketchField `json:"field,omitempty"`
	SplitPoints []float64                    `json:"splitPoints,omitempty"`
}

// NewQuantilesDoublesSketchToCDF new instance of QuantilesDoublesSketchToCDF
func NewQuantilesDoublesSketchToCDF() *QuantilesDoublesSketchToCDF {
	q := &QuantilesDoublesSketchToCDF{}
	q.SetType("quantilesDoublesSketchToCDF")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketchToCDF) SetName(name string) *QuantilesDoublesSketchToCDF {
	q.Base.SetName(name)
	return q
}

// SetSplitPoints set splitPoints
func (q *QuantilesDoublesSketchToCDF) SetSplitPoints(splitPoints []float64) *QuantilesDoublesSketchToCDF {
	q.SplitPoints = splitPoints
	return q
}

// SetField set QuantilesDoublesSketchField
func (q *QuantilesDoublesSketchToCDF) SetField(field *QuantilesDoublesSketchField) *QuantilesDoublesSketchToCDF {
	q.Field = field
	return q
}
