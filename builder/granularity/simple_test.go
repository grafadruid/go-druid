package granularity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSimple(t *testing.T) {
	s := NewSimple()
	s.SetGranularity("all")

	expected := `"all"`
	built, err := Load([]byte(expected))
	assert.Nil(t, err)

	assert.Equal(t, s, built, "expected and generated do not match")
}
