package aggregation

import (
	"encoding/json"
	"testing"

	"github.com/grafadruid/go-druid/builder/filter"
	"github.com/grafadruid/go-druid/builder/types"
	"github.com/stretchr/testify/assert"
)

func TestLoadFilteredAggregator(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte(`{
			  "type": "filtered",
			  "aggregator": {
				"type": "count",
				"name": "a1"
			  },
			  "filter": {
				"type": "bound",
				"dimension": "__time",
				"lower": "1629136173000",
				"upper": null,
				"lowerStrict": true,
				"upperStrict": false,
				"extractionFn": null,
				"ordering": "numeric"
			  },
			  "name": "a1"
			}`))
	assert.NotNil(f)
	assert.Nil(err, "error should be nil")
}

func TestNewFiltered(t *testing.T) {
	f := filter.NewBound().SetLower("1629136173000").SetLowerStrict(true).SetDimension("__time").SetOrdering(types.Numeric).
		SetUpperStrict(false)
	a := NewCount().SetName("a1")
	fa := NewFiltered().SetFilter(f).SetAggregator(a).SetName("a1")

	expected := `{
      "type": "filtered",
      "aggregator": {
        "type": "count",
        "name": "a1"
      },
      "filter": {
        "type": "bound",
        "dimension": "__time",
        "lower": "1629136173000",
        "lowerStrict": true,
        "upperStrict": false,
        "ordering": "numeric"
      },
      "name": "a1"
    }`

	faJSON, err := json.Marshal(fa)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(faJSON))
}
