package aggregation

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/dimension"
	"github.com/grafadruid/go-druid/builder/extractionfn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardinality(t *testing.T) {
	cardinality := NewCardinality()
	substringExtra := extractionfn.NewSubstring().SetIndex(0).SetLength(1)
	extraction := dimension.NewExtraction().SetDimension("last_name").
		SetOutputName("last_name_first_char").
		SetExtractionFn(substringExtra)

	cardinality.SetName("distinct_last_name_first_char").
		SetFields([]builder.DimensionSpec{extraction}).SetByRow(true).SetRound(false)
	expected := `{
  "type": "cardinality",
  "name": "distinct_last_name_first_char",
  "fields": [
    {
     "type" : "extraction",
     "dimension" : "last_name",
     "outputName" :  "last_name_first_char",
     "extractionFn" : { "type" : "substring", "index" : 0, "length" : 1 }
    }
  ],
  "byRow" : true,
  "round" : false
}`
	cardinalityJSON, err := json.Marshal(cardinality)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(cardinalityJSON))
}
