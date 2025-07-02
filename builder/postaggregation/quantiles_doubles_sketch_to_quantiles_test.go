package postaggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantilesDoublesSketchToQuantiles(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("tp75tp90").SetFieldName("a1:agg")
	quantilesDoublesSketchToQuantiles := NewQuantilesDoublesSketchToQuantiles()
	quantilesDoublesSketchToQuantiles.SetName("tp75tp90").SetField(qf).SetFractions([]float64{0.75, 0.90})

	// "omitempty" will ignore boolean=false
	quantilesDoublesSketchToQuantilesJSON := `
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

	t.Run("build quantilesDoublesSketchToQuantiles",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesDoublesSketchToQuantiles)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesDoublesSketchToQuantilesJSON)
		})

	t.Run("load quantilesDoublesSketchToQuantiles",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesDoublesSketchToQuantilesJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketchToQuantiles,
				postAgg)
		})
}
