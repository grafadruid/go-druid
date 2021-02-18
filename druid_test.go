package druid

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBaseURLWithSuffix(t *testing.T) {
	assert := assert.New(t)

	d, err := NewClient("localhost:8082")
	assert.Nil(err, "error should be nil")
	assert.NotNil(d, "client should not be nil")

	wantBaseURL, _ := url.ParseRequestURI("/")
	err = d.setBaseURL("")
	assert.Nil(err, "error should be nil")
	assert.Equal(d.baseURL, wantBaseURL, "they should not be equal")
}
