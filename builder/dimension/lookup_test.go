package dimension

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDimensionLookup(t *testing.T) {
	lookup := NewLookup().SetOutputName("output_name").SetName("lookup_name").SetRetainMissingValue(true)

	// "omitempty" will ignore boolean=false
	expected := `{"name":"lookup_name", "outputName":"output_name", "retainMissingValue":true, "type":"lookup"}`

	lookupJSON, err := json.Marshal(lookup)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(lookupJSON))
}
