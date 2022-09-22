package aggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantilesDoublesSketch(t *testing.T) {
	quantilesDoublesSketch := NewQuantilesDoublesSketch()
	quantilesDoublesSketch.SetName("output_name").SetFieldName("metric_name").SetK(100)
	// "omitempty" will ignore boolean=false
	quantilesDoublesSketchJSON := `{"type":"quantilesDoublesSketch", "name":"output_name", "fieldName": "metric_name", "k":100}`

	t.Run("build quantilesDoublesSketch",
		func(t *testing.T) {
			aggJSON, err := json.Marshal(quantilesDoublesSketch)
			assert.Nil(t,
				err)
			assert.JSONEq(t,
				quantilesDoublesSketchJSON,
				string(aggJSON))
		})

	t.Run("load quantilesDoublesSketch",
		func(t *testing.T) {
			agg, err := Load([]byte(quantilesDoublesSketchJSON))
			assert.Nil(t,
				err)
			assert.Equal(t,
				quantilesDoublesSketch,
				agg)
		})
}
