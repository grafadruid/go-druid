package postaggregation

// QuantilesDoublesSketchField struct for Field in QuantilesDoublesSketch Post Aggregators
type QuantilesDoublesSketchField struct {
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
}

// NewQuantilesDoublesSketchField new instance of QuantilesDoublesSketchField
func NewQuantilesDoublesSketchField() *QuantilesDoublesSketchField {
	return &QuantilesDoublesSketchField{}
}

// SetName set name
func (qf *QuantilesDoublesSketchField) SetName(name string) *QuantilesDoublesSketchField {
	qf.Name = name
	return qf
}

// SetType set type
func (qf *QuantilesDoublesSketchField) SetType(typ string) *QuantilesDoublesSketchField {
	qf.Type = typ
	return qf
}

// SetFieldName set fieldName
func (qf *QuantilesDoublesSketchField) SetFieldName(fieldName string) *QuantilesDoublesSketchField {
	qf.FieldName = fieldName
	return qf
}
