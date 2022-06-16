package postaggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantilesDoublesSketchToQuantile(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("tp90").SetFieldName("a1:agg")
	quantilesDoublesSketchToQuantile := NewQuantilesDoublesSketchToQuantile()
	quantilesDoublesSketchToQuantile.SetName("tp90").SetField(qf).SetFraction(0.90)

	// "omitempty" will ignore boolean=false
	expected := `
{
  "type": "quantilesDoublesSketchToQuantile",
  "name": "tp90",
  "field": {
    "type": "fieldAccess",
    "name": "tp90",
    "fieldName": "a1:agg"
  },
  "fraction": 0.9
}
`

	quantilesDoublesSketchToQuantileJSON, err := json.Marshal(quantilesDoublesSketchToQuantile)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(quantilesDoublesSketchToQuantileJSON))
}
