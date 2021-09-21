package limitspec

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	// test error unsupported limitSpec
	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported limitspec type")

	// set up good test case data
	var limitOrderByColumn = OrderByColumnSpec{
		Dimension:      "counter",
		Direction:      "descending",
		DimensionOrder: "numeric",
	}
	testLimitSpec := NewDefault().SetLimit(10).SetOffset(10).SetColumns([]OrderByColumnSpec{limitOrderByColumn})
	filterQuery := `{"type":"default","columns":[{"dimension":"counter","direction":"descending","dimensionOrder":"numeric"}],"offset":10,"limit":10}`
	limitSpecJson, err := json.Marshal(testLimitSpec)
	assert.Nil(err)
	assert.JSONEq(filterQuery, string(limitSpecJson))

	// test Load
	f, err = Load([]byte(filterQuery))
	assert.Nil(err)
	assert.NotNil(f)
	assert.Equal(testLimitSpec, f)

}