package havingspec

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported havingspec type")
}

func TestEqualTo(t *testing.T) {
	having := NewEqualTo()
	having.SetValue(0)
	having.SetAggregation("agg")

	got, err := json.Marshal(having)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"aggregation":"agg", "type":"equalTo", "value":0}`, string(got))
}
