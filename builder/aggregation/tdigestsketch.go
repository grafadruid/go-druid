package aggregation

// TDigestSketch holds the tdigest sketch struct based on
// Aggregator section in https://druid.apache.org/docs/latest/development/extensions-contrib/tdigestsketch-quantiles.html
type TDigestSketch struct {
	Base
	FieldName   string `json:"fieldName,omitempty"`
	Compression int64  `json:"compression,omitempty"`
}

// NewTDigestSketch create a new instance of TDigestSketch
func NewTDigestSketch() *TDigestSketch {
	t := &TDigestSketch{}
	t.Base.SetType("tDigestSketch")
	return t
}

// SetName set name
func (t *TDigestSketch) SetName(name string) *TDigestSketch {
	t.Base.SetName(name)
	return t
}

// SetFieldName set fieldName
func (t *TDigestSketch) SetFieldName(fieldName string) *TDigestSketch {
	t.FieldName = fieldName
	return t
}

// SetCompression set tdigest compression
func (t *TDigestSketch) SetCompression(compression int64) *TDigestSketch {
	t.Compression = compression
	return t
}
