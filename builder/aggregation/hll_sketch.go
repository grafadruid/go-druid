package aggregation

// HLLSketch holds the HLL sketch struct based on
// Aggregator section in https://druid.apache.org/docs/latest/development/extensions-core/datasketches-hll.html
type HLLSketch struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	LgK        int64  `json:"lgK,omitempty"`
	TgtHLLType string `json:"tgtHllType,omitempty"`
	Round      *bool  `json:"round,omitempty"`
}

// NewHLLSketchBuild create a new instance of HLLSketch with type HLLSketchBuild
func NewHLLSketchBuild() *HLLSketch {
	t := &HLLSketch{}
	t.Base.SetType("HLLSketchBuild")
	return t
}

// NewHLLSketchMerge create a new instance of HLLSketch with type HLLSketchMerge
func NewHLLSketchMerge() *HLLSketch {
	t := &HLLSketch{}
	t.Base.SetType("HLLSketchMerge")
	return t
}

// SetName set name
func (t *HLLSketch) SetName(name string) *HLLSketch {
	t.Base.SetName(name)
	return t
}

// SetFieldName set fieldName
func (t *HLLSketch) SetFieldName(fieldName string) *HLLSketch {
	t.FieldName = fieldName
	return t
}

// SetLgK set lgK. The value needs to be in the [4, 21] range
func (t *HLLSketch) SetLgK(lgk int64) *HLLSketch {
	t.LgK = lgk
	return t
}

// SetTgtHllType set tgtHllType
func (t *HLLSketch) SetTgtHLLType(tgtHLLType string) *HLLSketch {
	t.TgtHLLType = tgtHLLType
	return t
}

// SetRound set round.
func (t *HLLSketch) SetRound(round bool) *HLLSketch {
	t.Round = &round
	return t
}
