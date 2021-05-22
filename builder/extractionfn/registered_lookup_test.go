package extractionfn

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisteredLookupExtractionFn(t *testing.T) {
	lookup := NewRegisteredLookup()
	lookup.SetLookup("my_lookup").SetOptimize(false).SetReplaceMissingValueWith("N/A")

	// it's important that 'false' values are not omitted in the 'injective' and 'optimize' fields.
	expected := `{"type":"registeredLookup", "lookup":"my_lookup", "optimize":false, "replaceMissingValueWith":"N/A"}`

	lookupJSON, err := json.Marshal(lookup)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(lookupJSON))
}
