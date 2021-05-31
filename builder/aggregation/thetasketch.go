package aggregation

// ThetaSketch holds the theta sketch struct based on
// Aggregator section in http://druid.apache.org/docs/latest/development/extensions-core/datasketches-theta.html
type ThetaSketch struct {
	Base
	FieldName          string `json:"fieldName,omitempty"`
	IsInputThetaSketch *bool  `json:"isInputThetaSketch,omitempty"`
	Size               int64  `json:"size,omitempty"`
}

// NewThetaSketch create a new instance of ThetaSketch
func NewThetaSketch() *ThetaSketch {
	t := &ThetaSketch{}
	t.Base.SetType("thetaSketch")
	return t
}

// SetName set name
func (t *ThetaSketch) SetName(name string) *ThetaSketch {
	t.Base.SetName(name)
	return t
}

// SetFieldName set fieldName
func (t *ThetaSketch) SetFieldName(fieldName string) *ThetaSketch {
	t.FieldName = fieldName
	return t
}

// SetIsInputThetaSketch set theta isInputThetaSketch
func (t *ThetaSketch) SetIsInputThetaSketch(isInputThetaSketch bool) *ThetaSketch {
	t.IsInputThetaSketch = &isInputThetaSketch
	return t
}

// SetSize set theta size
func (t *ThetaSketch) SetSize(size int64) *ThetaSketch {
	t.Size = size
	return t
}
