package builder

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGeneric(t *testing.T) {
	generic := NewGeneric("customAggregation")

	expected := `{"type":"customAggregation"}`

	genericJSON, err := json.Marshal(generic)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(genericJSON))
}

func TestGeneric_SetField(t *testing.T) {
	t.Run("common field", func(t *testing.T) {
		withSpecificMethod := NewGeneric("customAggregation")
		withSpecificMethod.
			SetType("overriddenAggregation").
			SetName("op")

		withSetFieldMethod := NewGeneric("customAggregation").
			SetField("type", "overriddenAggregation").
			SetField("name", "op")

		assert.Equal(t, withSpecificMethod, withSetFieldMethod)

		expected := `{"type":"overriddenAggregation","name":"op"}`

		genericJSON, err := json.Marshal(withSetFieldMethod)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(genericJSON))
	})
	t.Run("any field", func(t *testing.T) {
		generic := NewGeneric("customAggregation").SetField("foo", "bar")

		expected := `{"type":"customAggregation","foo":"bar"}`

		genericJSON, err := json.Marshal(generic)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(genericJSON))
	})
}

func TestGeneric_MergeFields(t *testing.T) {
	generic := NewGeneric("customAggregation").
		SetField("one", "direct").
		SetField("two", "direct").
		MergeFields(map[string]interface{}{
			"two":   "merged",
			"three": "merged",
		})

	expected := `{"type":"customAggregation","one":"direct","two":"merged","three":"merged"}`

	genericJSON, err := json.Marshal(generic)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(genericJSON))
}

func TestGeneric_SetFields(t *testing.T) {
	generic := NewGeneric("customAggregation").
		SetField("one", "direct").
		SetField("two", "direct").
		SetFields(map[string]interface{}{
			"two":   "overridden",
			"three": "overridden",
		})

	expected := `{"type":"customAggregation","two":"overridden","three":"overridden"}`

	genericJSON, err := json.Marshal(generic)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(genericJSON))
}
