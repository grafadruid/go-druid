package dimension

import (
	"encoding/json"
	lookup2 "github.com/grafadruid/go-druid/builder/lookup"
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

func TestNewLookup(t *testing.T) {
	expected := `{
	 "type":"lookup",
	 "dimension":"dimensionName",
	 "outputName":"dimensionOutputName",
	 "replaceMissingValueWith":"missing_value",
	 "retainMissingValue":false,
	 "lookup":{"type": "map", "map":{"key":"value"}, "isOneToOne":false}
	}`
	mapFn := lookup2.NewMap().SetMap(map[string]string{"key": "value"}).SetIsOneToOne(false)
	lookup := NewLookup().SetOutputName("dimensionOutputName").SetRetainMissingValue(false).
		SetDimension("dimensionName").
		SetReplaceMissingValueWith("missing_value").SetLookup(mapFn)

	lookupJSON, err := json.Marshal(lookup)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(lookupJSON))
}
