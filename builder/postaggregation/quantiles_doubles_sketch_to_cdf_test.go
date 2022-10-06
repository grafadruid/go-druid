package postaggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantilesDoublesSketchToCDF(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("tp75tp90").SetFieldName("a1:agg")
	quantilesDoublesSketchToCDF := NewQuantilesDoublesSketchToCDF()
	quantilesDoublesSketchToCDF.SetName("tp75tp90").SetField(qf).SetSplitPoints([]float64{0.75, 0.90})

	// "omitempty" will ignore boolean=false
	quantilesDoublesSketchToCDFJSON := `
				{
				  "type": "quantilesDoublesSketchToCDF",
				  "name": "tp75tp90",
				  "field": {
					"type": "fieldAccess",
					"name": "tp75tp90",
					"fieldName": "a1:agg"
				  },
				  "splitPoints": [0.75, 0.9]
				}
				`

	t.Run("build quantilesDoublesSketchToQuantiles",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesDoublesSketchToCDF)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesDoublesSketchToCDFJSON)
		})

	t.Run("load quantilesDoublesSketchToQuantiles",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesDoublesSketchToCDFJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketchToCDF,
				postAgg)
		})
}
