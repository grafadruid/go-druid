package testutil

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Compare compares the builder type b to expected and also built to expected.
// json(b) == json(built) == expected
func Compare(t *testing.T, expected []byte, b interface{}, built interface{}) {
	// convert builder b to JSON so we can compare the JSON of builder to expected JSON.
	js, err := json.Marshal(b)
	assert.Nil(t, err)

	assert.Equal(t, js, expected)

	assert.Equal(t, b, built)

	// convert built (which is generated from expected JSON) to JSON so it can also
	// be compared with expected.
	jbuilt, err := json.Marshal(built)
	assert.Nil(t, err)
	assert.Equal(t, jbuilt, expected)
}
