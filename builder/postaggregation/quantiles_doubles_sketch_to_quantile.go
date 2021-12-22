package postaggregation

// QuantilesDoublesSketchToQuantile struct based on
// PostAggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html#quantile
type QuantilesDoublesSketchToQuantile struct {
	Base
	Field    *QuantilesDoublesSketchToQuantileField `json:"field,omitempty"`
	Fraction float64                                `json:"fraction,omitempty"`
}

// QuantilesDoublesSketchToQuantileField struct for Field in QuantilesDoublesSketchToQuantile
type QuantilesDoublesSketchToQuantileField struct {
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
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
	q.Fraction = fraction
	return q
}

// SetField set QuantilesDoublesSketchToQuantile
func (q *QuantilesDoublesSketchToQuantile) SetField(field *QuantilesDoublesSketchToQuantileField) *QuantilesDoublesSketchToQuantile {
	q.Field = field
	return q
}

func NewQuantilesDoublesSketchToQuantileField() *QuantilesDoublesSketchToQuantileField {
	qf := &QuantilesDoublesSketchToQuantileField{}
	return qf
}

// SetName set name
func (qf *QuantilesDoublesSketchToQuantileField) SetName(name string) *QuantilesDoublesSketchToQuantileField{
	qf.Name = name
	return qf
}

// SetType set type
func (qf *QuantilesDoublesSketchToQuantileField) SetType(typ string) *QuantilesDoublesSketchToQuantileField {
	qf.Type = typ
	return qf
}

// SetFieldName set fieldName
func (qf *QuantilesDoublesSketchToQuantileField) SetFieldName(fieldName string) *QuantilesDoublesSketchToQuantileField {
	qf.FieldName = fieldName
	return qf
}
