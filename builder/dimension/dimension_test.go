package dimension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported dimension type")
}

func TestLoadExtractionDimension(t *testing.T) {

	f, err := Load([]byte(`{
      "type": "extraction",
      "dimension": "lookupKey",
      "outputName": "assetTier",
      "extractionFn": {
        "type": "cascade",
        "extractionFns": [
          {
            "type": "registeredLookup",
            "lookup": "asset_id_to_metadata",
            "retainMissingValue": true
          },
          {
            "type": "regex",
            "expr": "(?<=tier\":\")(.*?)(?=\")",
            "replaceMissingValue": false
          }
        ]
      }
    }`))
	assert.Nil(t, err)
	assert.NotNil(t, f)
}
