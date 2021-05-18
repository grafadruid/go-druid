package aggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHLLSketchBuild(t *testing.T) {
	hllSketch := NewHLLSketchBuild()
	hllSketch.SetName("output_name").SetFieldName("metric_name").SetLgK(5).SetTgtHllType("HLL_6")

	// "omitempty" will ignore boolean=false
	expected := `{"type":"HLLSketchBuild", "name":"output_name", "fieldName": "metric_name", "lgK":5, "tgtHllType":"HLL_6"}`

	hllSketchJson, err := json.Marshal(hllSketch)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(hllSketchJson))
}

func TestHLLSketchMerge(t *testing.T) {
	hllSketch := NewHLLSketchMerge()
	hllSketch.SetName("output_name").SetFieldName("metric_name").SetLgK(5).SetTgtHllType("HLL_6")

	// "omitempty" will ignore boolean=false
	expected := `{"type":"HLLSketchMerge", "name":"output_name",  "fieldName": "metric_name", "lgK":5, "tgtHllType":"HLL_6"}`

	hllSketchJson, err := json.Marshal(hllSketch)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(hllSketchJson))
}
