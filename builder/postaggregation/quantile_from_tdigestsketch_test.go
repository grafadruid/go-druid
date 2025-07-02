package postaggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantileFromTDigestSketch(t *testing.T) {
	qf := NewQuantileFromTDigestSketchField()
	qf.SetType("fieldAccess").SetFieldName("merged_sketch")
	quantilesFromTDigestSketch := NewQuantileFromTDigestSketch()
	quantilesFromTDigestSketch.SetName("tp90").SetField(qf).SetFraction(0.90)

	// "omitempty" will ignore boolean=false
	quantileFromTDigestSketchJSON := `
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
	t.Run("build quantileFromTDigestSketch",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesFromTDigestSketch)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				quantileFromTDigestSketchJSON,
				string(postAggJSON))
		})

	t.Run("load quantileFromTDigestSketch",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantileFromTDigestSketchJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesFromTDigestSketch,
				postAgg)
		})
}
