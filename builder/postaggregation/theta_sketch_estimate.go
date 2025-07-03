package postaggregation

// ThetaSketchEstimate is a post-aggregation that estimates the size of a theta sketch.
// https://druid.apache.org/docs/latest/development/extensions-core/datasketches-theta#sketch-estimator
type ThetaSketchEstimate struct {
	Base
	Field map[string]string `json:"field,omitempty"`
}

// NewThetaSketchEstimate creates a new instance of ThetaSketchEstimate.
func NewThetaSketchEstimate() *ThetaSketchEstimate {
	t := &ThetaSketchEstimate{}
	t.SetType("thetaSketchEstimate")

	return t
}

// SetName sets the name of the ThetaSketchEstimate.
func (t *ThetaSketchEstimate) SetName(name string) *ThetaSketchEstimate {
	t.Base.SetName(name)

	return t
}

// SetField sets the field for the ThetaSketchEstimate.
func (t *ThetaSketchEstimate) SetField(fieldName string) *ThetaSketchEstimate {
	t.Field = map[string]string{
		"type":      "fieldAccess",
		"fieldName": fieldName,
	}

	return t
}
