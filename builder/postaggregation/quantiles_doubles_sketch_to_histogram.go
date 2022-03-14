package postaggregation

// QuantilesDoublesSketchToHistogram struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html#histogram
type QuantilesDoublesSketchToHistogram struct {
	Base
	Field       *QuantilesDoublesSketchField `json:"field,omitempty"`
	SplitPoints []float64                    `json:"splitPoints,omitempty"`
	NumBins     int64                        `json:"numBins,omitempty"`
}

// NewQuantilesDoublesSketchToHistogram new instance of QuantilesDoublesSketchToHistogram
func NewQuantilesDoublesSketchToHistogram() *QuantilesDoublesSketchToHistogram {
	q := &QuantilesDoublesSketchToHistogram{}
	q.SetType("quantilesDoublesSketchToHistogram")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketchToHistogram) SetName(name string) *QuantilesDoublesSketchToHistogram {
	q.Base.SetName(name)
	return q
}

// SetSplitPoints set splitPoints
func (q *QuantilesDoublesSketchToHistogram) SetSplitPoints(splitPoints []float64) *QuantilesDoublesSketchToHistogram {
	q.SplitPoints = splitPoints
	return q
}

// SetNumBins set namBins
func (q *QuantilesDoublesSketchToHistogram) SetNumBins(numBins int64) *QuantilesDoublesSketchToHistogram {
	q.NumBins = numBins
	return q
}

// SetField set QuantilesDoublesSketchField
func (q *QuantilesDoublesSketchToHistogram) SetField(field *QuantilesDoublesSketchField) *QuantilesDoublesSketchToHistogram {
	q.Field = field
	return q
}
