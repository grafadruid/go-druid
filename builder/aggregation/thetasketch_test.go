package aggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThetaSketch(t *testing.T) {
	thetaSketch := NewThetaSketch()
	thetaSketch.SetName("output_name").SetFieldName("metric_name").SetIsInputThetaSketch(false).SetSize(16384)

	// "omitempty" will ignore boolean=false
	expected := `{"type":"thetaSketch", "name":"output_name", "fieldName":"metric_name", "size":16384}`

	thetaSketchJSON, err := json.Marshal(thetaSketch)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(thetaSketchJSON))
}
