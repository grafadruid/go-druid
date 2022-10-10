package postaggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantilesDoublesSketchToString(t *testing.T) {
	qf := NewQuantilesDoublesSketchField()
	qf.SetType("fieldAccess").SetName("tp75tp90").SetFieldName("a1:agg")
	quantilesDoublesSketchToString := NewQuantilesDoublesSketchToString()
	quantilesDoublesSketchToString.SetName("tp75tp90").SetField(qf)

	// "omitempty" will ignore boolean=false
	quantilesDoublesSketchToStringJSON := `
				{
				  "type": "quantilesDoublesSketchToString",
				  "name": "tp75tp90",
				  "field": {
					"type": "fieldAccess",
					"name": "tp75tp90",
					"fieldName": "a1:agg"
				  }
				}
				`

	t.Run("build quantilesDoublesSketchToString",
		func(t *testing.T) {
			postAggJSON, err := json.Marshal(quantilesDoublesSketchToString)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				string(postAggJSON),
				quantilesDoublesSketchToStringJSON)
		})

	t.Run("load quantilesDoublesSketchToString",
		func(t *testing.T) {
			postAgg, err := Load([]byte(quantilesDoublesSketchToStringJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketchToString,
				postAgg)
		})
}
