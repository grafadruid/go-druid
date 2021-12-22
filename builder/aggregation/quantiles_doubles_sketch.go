package aggregation


// QuantilesDoublesSketch holds the Quantiles Doubles sketch based on
// https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html
type QuantilesDoublesSketch struct {
	Base
	FieldName string `json:"fieldName, omitempty"`
	K         int64  `json:"k, omitempty"`
}

// NewQuantilesDoublesSketch create a new instance of QuantilesDoublesSketch with type QuantilesDoublesSketch
func NewQuantilesDoublesSketch() *QuantilesDoublesSketch {
	t := &QuantilesDoublesSketch{}
	t.Base.SetType("quantilesDoublesSketch")
	return t
}

// SetName set name
func (t *QuantilesDoublesSketch) SetName(name string) *QuantilesDoublesSketch {
	t.Base.SetName(name)
	return t
}

// SetFieldName set fieldName
func (t *QuantilesDoublesSketch) SetFieldName(fieldName string) *QuantilesDoublesSketch {
	t.FieldName = fieldName
	return t
}

// SetK set K. The value needs to must be a power of 2 from 2 to 32768
func (t *QuantilesDoublesSketch) SetK(k int64) *QuantilesDoublesSketch {
	t.K = k
	return t
}
