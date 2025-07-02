package postaggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantilesDoublesSketchToHistogram_SetNumBins(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("h").SetFieldName("a1:agg")
	quantilesDoublesSketchToHistogram := NewQuantilesDoublesSketchToHistogram()
	quantilesDoublesSketchToHistogram.SetName("h").SetField(qf).SetNumBins(2)

	// "omitempty" will ignore boolean=false
	quantilesDoublesSketchToHistogramJSON := `
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

	t.Run("build quantilesDoublesSketchToHistogram",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesDoublesSketchToHistogram)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesDoublesSketchToHistogramJSON)
		})

	t.Run("load quantilesDoublesSketchToHistogram",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesDoublesSketchToHistogramJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketchToHistogram,
				postAgg)
		})
}

func TestQuantilesDoublesSketchToHistogram_SetSplitPoints(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("h").SetFieldName("a1:agg")
	quantilesDoublesSketchToHistogram := NewQuantilesDoublesSketchToHistogram()
	quantilesDoublesSketchToHistogram.SetName("h").SetField(qf).SetSplitPoints([]float64{0.5, 1.5, 2.0})

	// "omitempty" will ignore boolean=false
	quantilesDoublesSketchToHistogramJSON := `
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

	t.Run("build quantilesDoublesSketchToHistogram",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesDoublesSketchToHistogram)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesDoublesSketchToHistogramJSON)
		})

	t.Run("load quantilesDoublesSketchToHistogram",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesDoublesSketchToHistogramJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketchToHistogram,
				postAgg)
		})
}
