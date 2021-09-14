package virtualcolumn

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExpression(t *testing.T) {
	expression := NewExpression()
	expression.SetName("dim1").SetExpression("lower(\"dim1\")").SetOutputType(types.String)

	expected := `{"type":"expression", "name":"dim1", "expression": "lower(\"dim1\")", "outputType":"STRING"}`

	expressionJSON, err := json.Marshal(expression)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(expressionJSON))
}
