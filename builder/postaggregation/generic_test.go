package postaggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGeneric(t *testing.T) {
	generic := NewGeneric("customPostAggregation")

	expected := `{"type":"customPostAggregation"}`

	genericJSON, err := json.Marshal(generic)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(genericJSON))
}

func TestGeneric_SetField(t *testing.T) {
	t.Run("common field", func(t *testing.T) {
		withSpecificMethod := NewGeneric("customPostAggregation")
		withSpecificMethod.
			SetType("overriddenPostAggregation").
			SetName("op")

		withSetFieldMethod := NewGeneric("customPostAggregation").
			SetField("type", "overriddenPostAggregation").
			SetField("name", "op")

		assert.Equal(t, withSpecificMethod, withSetFieldMethod)

		expected := `{"type":"overriddenPostAggregation","name":"op"}`

		genericJSON, err := json.Marshal(withSetFieldMethod)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(genericJSON))
	})
	t.Run("any field", func(t *testing.T) {
		generic := NewGeneric("customPostAggregation").SetField("foo", "bar")

		expected := `{"type":"customPostAggregation","foo":"bar"}`

		genericJSON, err := json.Marshal(generic)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(genericJSON))
	})
}

func TestGeneric_MergeFields(t *testing.T) {
	generic := NewGeneric("customPostAggregation").
		SetField("one", "direct").
		SetField("two", "direct").
		MergeFields(map[string]interface{}{
			"two":   "merged",
			"three": "merged",
		})

	expected := `{"type":"customPostAggregation","one":"direct","two":"merged","three":"merged"}`

	genericJSON, err := json.Marshal(generic)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(genericJSON))
}

func TestGeneric_SetFields(t *testing.T) {
	generic := NewGeneric("customPostAggregation").
		SetField("one", "direct").
		SetField("two", "direct").
		SetFields(map[string]interface{}{
			"two":   "overridden",
			"three": "overridden",
		})

	expected := `{"type":"customPostAggregation","two":"overridden","three":"overridden"}`

	genericJSON, err := json.Marshal(generic)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(genericJSON))
}
