package aggregation

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJavascript(t *testing.T) {
	javaScript := NewJavascript()
	javaScript.SetName("sum(log(x)*y) + 10").
		SetFieldNames([]string{"x", "y"}).
		SetFnAggregate("function(current, a, b)      { return current + (Math.log(a) * b); }").
		SetFnCombine("function(partialA, partialB) { return partialA + partialB; }").
		SetFnReset("function()                   { return 10; }")
	expected := `{
  "type": "javascript",
  "name": "sum(log(x)*y) + 10",
  "fieldNames": ["x", "y"],
  "fnAggregate" : "function(current, a, b)      { return current + (Math.log(a) * b); }",
  "fnCombine"   : "function(partialA, partialB) { return partialA + partialB; }",
  "fnReset"     : "function()                   { return 10; }"
}`
	javaScriptJSON, err := json.Marshal(javaScript)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(javaScriptJSON))
}
