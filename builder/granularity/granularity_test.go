package granularity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte(`{"type": "_not_such_Type"}`))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported granularity type")
}
