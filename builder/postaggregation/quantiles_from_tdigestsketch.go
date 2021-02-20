package postaggregation

// QuantilesFromTDigestSketch struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-contrib/tdigestsketch-quantiles.html
type QuantilesFromTDigestSketch struct {
	Base
	Fractions []float64                        `json:"fractions,omitempty"`
	Field     *QuantilesFromTDigestSketchField `json:"field,omitempty"`
}

// QuantilesFromTDigestSketchField struct for Field in QuantilesFromTDigestSketch
type QuantilesFromTDigestSketchField struct {
	Type      string `json:"type,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
}

// NewQuantilesFromTDigestSketch new instance of QuantilesFromTDigestSketch
func NewQuantilesFromTDigestSketch() *QuantilesFromTDigestSketch {
	q := &QuantilesFromTDigestSketch{}
	q.SetType("quantilesFromTDigestSketch")
	return q
}

// SetName set name
func (q *QuantilesFromTDigestSketch) SetName(name string) *QuantilesFromTDigestSketch {
	q.Base.SetName(name)
	return q
}

// SetFractions set fractions
func (q *QuantilesFromTDigestSketch) SetFractions(fractions []float64) *QuantilesFromTDigestSketch {
	q.Fractions = fractions
	return q
}

// SetField set QuantilesFromTDigestSketchField
func (q *QuantilesFromTDigestSketch) SetField(field *QuantilesFromTDigestSketchField) *QuantilesFromTDigestSketch {
	q.Field = field
	return q
}

// NewQuantilesFromTDigestSketchField new instance of QuantilesFromTDigestSketchField
func NewQuantilesFromTDigestSketchField() *QuantilesFromTDigestSketchField {
	qf := &QuantilesFromTDigestSketchField{}
	return qf
}

// SetType set type
func (qf *QuantilesFromTDigestSketchField) SetType(typ string) *QuantilesFromTDigestSketchField {
	qf.Type = typ
	return qf
}

// SetFieldName set fieldName
func (qf *QuantilesFromTDigestSketchField) SetFieldName(fieldName string) *QuantilesFromTDigestSketchField {
	qf.FieldName = fieldName
	return qf
}
