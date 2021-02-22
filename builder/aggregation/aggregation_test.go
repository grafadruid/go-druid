package aggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported aggregation type")
}
