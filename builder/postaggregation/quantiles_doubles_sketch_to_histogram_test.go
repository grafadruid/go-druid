package postaggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantilesDoublesSketchToHistogram_SetNumBins(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("h").SetFieldName("a1:agg")
	quantilesDoublesSketchToHistogram := NewQuantilesDoublesSketchToHistogram()
	quantilesDoublesSketchToHistogram.SetName("h").SetField(qf).SetNumBins(2)

	// "omitempty" will ignore boolean=false
	expected := `
{
  "type": "quantilesDoublesSketchToHistogram",
  "name": "h",
  "numBins": 2,
  "field": {
    "type": "fieldAccess",
    "name": "h",
    "fieldName": "a1:agg"
  }
}
`

	quantilesDoublesSketchToHistogramJSON, err := json.Marshal(quantilesDoublesSketchToHistogram)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(quantilesDoublesSketchToHistogramJSON))
}

func TestQuantilesDoublesSketchToHistogram_SetSplitPoints(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("h").SetFieldName("a1:agg")
	quantilesDoublesSketchToHistogram := NewQuantilesDoublesSketchToHistogram()
	quantilesDoublesSketchToHistogram.SetName("h").SetField(qf).SetSplitPoints([]float64{0.5, 1.5, 2.0})

	// "omitempty" will ignore boolean=false
	expected := `
{
  "type": "quantilesDoublesSketchToHistogram",
  "name": "h",
  "splitPoints": [0.5,1.5,2.0],
  "field": {
    "type": "fieldAccess",
    "name": "h",
    "fieldName": "a1:agg"
  }
}
`

	quantilesDoublesSketchToHistogramJSON, err := json.Marshal(quantilesDoublesSketchToHistogram)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(quantilesDoublesSketchToHistogramJSON))
}
