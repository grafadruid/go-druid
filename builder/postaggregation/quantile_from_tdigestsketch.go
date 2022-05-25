package postaggregation

// QuantileFromTDigestSketch struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-contrib/tdigestsketch-quantiles.html
// See the "Similar to quantilesFromTDigestSketch except it takes in a single fraction for computing quantile" section
type QuantileFromTDigestSketch struct {
	Base
	Fraction *float64                        `json:"fraction,omitempty"`
	Field    *QuantileFromTDigestSketchField `json:"field,omitempty"`
}

// QuantileFromTDigestSketchField struct for Field in QuantileFromTDigestSketch
type QuantileFromTDigestSketchField struct {
	Type      string `json:"type,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
}

// NewQuantileFromTDigestSketch new instance of QuantileFromTDigestSketch
func NewQuantileFromTDigestSketch() *QuantileFromTDigestSketch {
	q := &QuantileFromTDigestSketch{}
	q.SetType("quantileFromTDigestSketch")
	return q
}

// SetName set name
func (q *QuantileFromTDigestSketch) SetName(name string) *QuantileFromTDigestSketch {
	q.Base.SetName(name)
	return q
}

// SetFraction set fraction
func (q *QuantileFromTDigestSketch) SetFraction(fraction float64) *QuantileFromTDigestSketch {
	q.Fraction = &fraction
	return q
}

// SetField set QuantileFromTDigestSketchField
func (q *QuantileFromTDigestSketch) SetField(field *QuantileFromTDigestSketchField) *QuantileFromTDigestSketch {
	q.Field = field
	return q
}

// NewQuantileFromTDigestSketchField new instance of QuantileFromTDigestSketchField
func NewQuantileFromTDigestSketchField() *QuantileFromTDigestSketchField {
	qf := &QuantileFromTDigestSketchField{}
	return qf
}

// SetType set type
func (qf *QuantileFromTDigestSketchField) SetType(typ string) *QuantileFromTDigestSketchField {
	qf.Type = typ
	return qf
}

// SetFieldName set fieldName
func (qf *QuantileFromTDigestSketchField) SetFieldName(fieldName string) *QuantileFromTDigestSketchField {
	qf.FieldName = fieldName
	return qf
}
