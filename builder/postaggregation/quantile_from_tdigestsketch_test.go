package postaggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantileFromTDigestSketch(t *testing.T) {
	qf := NewQuantileFromTDigestSketchField()
	qf.SetType("fieldAccess").SetFieldName("merged_sketch")
	quantilesFromTDigestSketch := NewQuantileFromTDigestSketch()
	quantilesFromTDigestSketch.SetName("tp90").SetField(qf).SetFraction(0.90)

	// "omitempty" will ignore boolean=false
	expected := `
{
  "type": "quantileFromTDigestSketch",
  "name": "tp90",
  "field": {
    "type": "fieldAccess",
    "fieldName": "merged_sketch"
  },
  "fraction": 0.9
}
`

	quantileFromTDigestSketchJSON, err := json.Marshal(quantilesFromTDigestSketch)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(quantileFromTDigestSketchJSON))
}
