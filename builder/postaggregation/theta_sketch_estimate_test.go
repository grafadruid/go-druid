package postaggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThetaSketchEstimate(t *testing.T) {
	ts := NewThetaSketchEstimate().
		SetName("testThetaSketchEstimate").
		SetField("testFiledSketch")
	raw := `
		{
		 "type": "thetaSketchEstimate",
		 "name": "testThetaSketchEstimate",
		 "field": {
		   "type": "fieldAccess",
		   "fieldName": "testFiledSketch"
		 }
		}
	`

	t.Run("build theta sketch estimate", func(t *testing.T) {
		got, err := json.Marshal(ts)
		assert.NoError(t, err)
		assert.JSONEq(t, string(got), raw)
	})

	t.Run("load theta sketch estimate", func(t *testing.T) {
		got, err := Load([]byte(raw))
		assert.NoError(t, err)
		assert.Equal(t, got, ts)
	})

	t.Run("round trip theta sketch estimate", func(t *testing.T) {
		gotRaw, err := json.Marshal(ts)
		assert.NoError(t, err)
		assert.JSONEq(t, string(gotRaw), raw)

		got, err := Load(gotRaw)
		assert.NoError(t, err)
		assert.Equal(t, got, ts)
	})
}
