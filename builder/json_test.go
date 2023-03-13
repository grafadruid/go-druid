package builder

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJSON(t *testing.T) {
	component := NewJSON("customComponent")

	expected := `{"type":"customComponent"}`

	componentJSON, err := json.Marshal(component)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(componentJSON))
}

func TestJSON_SetField(t *testing.T) {
	t.Run("common field", func(t *testing.T) {
		withSpecificMethod := NewJSON("customComponent")
		withSpecificMethod.
			SetType("overriddenComponent").
			SetName("op")

		withSetFieldMethod := NewJSON("customComponent").
			SetField("type", "overriddenComponent").
			SetField("name", "op")

		assert.Equal(t, withSpecificMethod, withSetFieldMethod)

		expected := `{"type":"overriddenComponent","name":"op"}`

		componentJSON, err := json.Marshal(withSetFieldMethod)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(componentJSON))
	})
	t.Run("any field", func(t *testing.T) {
		component := NewJSON("customComponent").SetField("foo", "bar")

		expected := `{"type":"customComponent","foo":"bar"}`

		componentJSON, err := json.Marshal(component)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(componentJSON))
	})
}

func TestJSON_MergeFields(t *testing.T) {
	component := NewJSON("customComponent").
		SetField("one", "direct").
		SetField("two", "direct").
		MergeFields(map[string]interface{}{
			"two":   "merged",
			"three": "merged",
		})

	expected := `{"type":"customComponent","one":"direct","two":"merged","three":"merged"}`

	componentJSON, err := json.Marshal(component)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(componentJSON))
}

func TestJSON_SetFields(t *testing.T) {
	component := NewJSON("customComponent").
		SetField("one", "direct").
		SetField("two", "direct").
		SetFields(map[string]interface{}{
			"two":   "overridden",
			"three": "overridden",
		})

	expected := `{"type":"customComponent","two":"overridden","three":"overridden"}`

	componentJSON, err := json.Marshal(component)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(componentJSON))
}
