package builder

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSpec(t *testing.T) {
	component := NewSpec("customComponent")

	expected := `{"type":"customComponent"}`

	componentJSON, err := json.Marshal(component)

	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(componentJSON))
}

func TestSpec_SetField(t *testing.T) {
	t.Run("common field", func(t *testing.T) {
		withSpecificMethod := NewSpec("customComponent")
		withSpecificMethod.
			SetType("overriddenComponent").
			SetName("op")

		withSetFieldMethod := NewSpec("customComponent").
			SetField("type", "overriddenComponent").
			SetField("name", "op")

		assert.Equal(t, withSpecificMethod, withSetFieldMethod)

		expected := `{"type":"overriddenComponent","name":"op"}`

		componentJSON, err := json.Marshal(withSetFieldMethod)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(componentJSON))
	})
	t.Run("any field", func(t *testing.T) {
		component := NewSpec("customComponent").SetField("foo", "bar")

		expected := `{"type":"customComponent","foo":"bar"}`

		componentJSON, err := json.Marshal(component)

		assert.NoError(t, err)
		assert.JSONEq(t, expected, string(componentJSON))
	})
}

func TestSpec_MergeFields(t *testing.T) {
	component := NewSpec("customComponent").
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

func TestSpec_SetFields(t *testing.T) {
	component := NewSpec("customComponent").
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
