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
	expected := `{"type":"quantilesDoublesSketch", "name":"output_name", "fieldName": "metric_name", "k":100}`

	quantilesDoublesSketchJSON, err := json.Marshal(quantilesDoublesSketch)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(quantilesDoublesSketchJSON))
}
