package limitspec

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported limitspec type")

	var limitOrderByColumn = OrderByColumnSpec{
		Dimension:      "counter",
		Direction:      "descending",
		DimensionOrder: "numeric",
	}
	testLimitSpec := NewDefault().SetLimit(10).SetOffset(10).SetColumns([]OrderByColumnSpec{limitOrderByColumn})

	expected := `{"type":"default","columns":[{"dimension":"counter","direction":"descending","dimensionOrder":"numeric"}],"offset":10,"limit":10}`

	limitSpecJson, err := json.Marshal(testLimitSpec)
	assert.Nil(err)
	fmt.Println(string(limitSpecJson))
	assert.JSONEq(expected, string(limitSpecJson))

}
