package dimension

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder/extractionfn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExtraction(t *testing.T) {
	substringExtra := extractionfn.NewSubstring().SetIndex(0).SetLength(1)
	extraction := NewExtraction().SetDimension("last_name").
		SetOutputName("last_name_first_char").
		SetExtractionFn(substringExtra)
	expected := `{
  "type": "extraction",
  "dimension": "last_name",
  "outputName": "last_name_first_char",
  "extractionFn": {
    "type": "substring",
    "index": 0,
    "length": 1
  }
}`
	extractionJSON, err := json.Marshal(extraction)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(extractionJSON))

}
