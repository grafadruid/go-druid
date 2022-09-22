package postaggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantilesFromTDigestSketch(t *testing.T) {
	qf := NewQuantilesFromTDigestSketchField()
	qf.SetType("fieldAccess").SetFieldName("merged_sketch")
	quantilesFromTDigestSketch := NewQuantilesFromTDigestSketch()
	quantilesFromTDigestSketch.SetName("tp75tp90").SetField(qf).SetFractions([]float64{0.75, 0.90})

	// "omitempty" will ignore boolean=false
	quantilesFromTDigestSketchJSON := `
{
  "type": "quantilesFromTDigestSketch",
  "name": "tp75tp90",
  "field": {
    "type": "fieldAccess",
    "fieldName": "merged_sketch"
  },
  "fractions": [0.75, 0.9]
}
`

	t.Run("build quantilesFromTDigestSketch",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesFromTDigestSketch)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesFromTDigestSketchJSON)
		})

	t.Run("load quantilesFromTDigestSketch",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesFromTDigestSketchJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesFromTDigestSketch,
				postAgg)
		})
}
