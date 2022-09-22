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
	quantilesDoublesSketchToQuantileJSON := `
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

	t.Run("build quantilesDoublesSketchToQuantile",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesDoublesSketchToQuantile)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesDoublesSketchToQuantileJSON)
		})

	t.Run("load quantilesDoublesSketchToQuantile",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesDoublesSketchToQuantileJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketchToQuantile,
				postAgg)
		})
}
