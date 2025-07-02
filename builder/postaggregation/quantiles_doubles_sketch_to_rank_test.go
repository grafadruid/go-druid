package postaggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantilesDoublesSketchToRank(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("tp90").SetFieldName("a1:agg")
	quantilesDoublesSketchToRank := NewQuantilesDoublesSketchToRank()
	quantilesDoublesSketchToRank.SetName("tp90").SetField(qf).SetValue(0.90)

	// "omitempty" will ignore boolean=false
	quantilesDoublesSketchToRankJSON := `
{
  "type": "quantilesDoublesSketchToRank",
  "name": "tp90",
  "field": {
    "type": "fieldAccess",
    "name": "tp90",
    "fieldName": "a1:agg"
  },
  "value": 0.9
}
`

	t.Run("build quantilesDoublesSketchToRank",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesDoublesSketchToRank)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesDoublesSketchToRankJSON)
		})

	t.Run("load quantilesDoublesSketchToRank",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesDoublesSketchToRankJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketchToRank,
				postAgg)
		})
}
