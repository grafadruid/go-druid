package aggregation

// QuantilesDoublesSketch holds the Quantiles Doubles sketch based on
// https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html
type QuantilesDoublesSketch struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
	K         int64  `json:"k,omitempty"`
}

// NewQuantilesDoublesSketch create a new instance of QuantilesDoublesSketch with type QuantilesDoublesSketch
func NewQuantilesDoublesSketch() *QuantilesDoublesSketch {
	q := &QuantilesDoublesSketch{}
	q.Base.SetType("quantilesDoublesSketch")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketch) SetName(name string) *QuantilesDoublesSketch {
	q.Base.SetName(name)
	return q
}

// SetFieldName set fieldName
func (q *QuantilesDoublesSketch) SetFieldName(fieldName string) *QuantilesDoublesSketch {
	q.FieldName = fieldName
	return q
}

// SetK set K. The value needs to must be a power of 2 from 2 to 32768
func (q *QuantilesDoublesSketch) SetK(k int64) *QuantilesDoublesSketch {
	q.K = k
	return q
}
