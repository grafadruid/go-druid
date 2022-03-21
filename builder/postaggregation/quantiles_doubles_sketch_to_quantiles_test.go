package postaggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantilesDoublesSketchToQuantiles(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("tp75tp90").SetFieldName("a1:agg")
	quantilesDoublesSketchToQuantiles := NewQuantilesDoublesSketchToQuantiles()
	quantilesDoublesSketchToQuantiles.SetName("tp75tp90").SetField(qf).SetFractions([]float64{0.75, 0.90})

	// "omitempty" will ignore boolean=false
	expected := `
{
  "type": "quantilesDoublesSketchToQuantiles",
  "name": "tp75tp90",
  "field": {
    "type": "fieldAccess",
    "name": "tp75tp90",
    "fieldName": "a1:agg"
  },
  "fractions": [0.75, 0.9]
}
`

	quantilesDoublesSketchToQuantilesJSON, err := json.Marshal(quantilesDoublesSketchToQuantiles)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(quantilesDoublesSketchToQuantilesJSON))
}
