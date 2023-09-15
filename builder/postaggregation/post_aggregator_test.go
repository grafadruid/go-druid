package postaggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h2oai/go-druid/builder"
)

func TestLoadInvalidJSON(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("not JSON value"))

	assert.Nilf(f, "post aggregation should be nil")
	assert.Errorf(err, "loading should return a non-nil error")
}

func TestLoadUntypedJSON(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{}"))

	assert.Nilf(f, "post aggregation should be nil")
	assert.Errorf(err, "loading should return a non-nil error")
}

func TestLoadValidAggregation(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected builder.Aggregator
	}{
		{
			name:     "arithmetic",
			jsonData: `{"type":"arithmetic"}`,
			expected: NewArithmetic().SetFields([]builder.PostAggregator{}),
		},
		{
			name:     "constant",
			jsonData: `{"type":"constant"}`,
			expected: NewConstant(),
		},
		{
			name:     "doubleGreatest",
			jsonData: `{"type":"doubleGreatest"}`,
			expected: NewDoubleGreatest().SetFields([]builder.PostAggregator{}),
		},
		{
			name:     "doubleLeast",
			jsonData: `{"type":"doubleLeast"}`,
			expected: NewDoubleLeast().SetFields([]builder.PostAggregator{}),
		},
		{
			name:     "expression",
			jsonData: `{"type":"expression"}`,
			expected: NewExpression(),
		},
		{
			name:     "fieldAccess",
			jsonData: `{"type":"fieldAccess"}`,
			expected: NewFieldAccess(),
		},
		{
			name:     "finalizingFieldAccess",
			jsonData: `{"type":"finalizingFieldAccess"}`,
			expected: NewFinalizingFieldAccess(),
		},
		{
			name:     "hyperUniqueFinalizing",
			jsonData: `{"type":"hyperUniqueFinalizing"}`,
			expected: NewHyperUniqueFinalizing(),
		},
		{
			name:     "javascript",
			jsonData: `{"type":"javascript"}`,
			expected: NewJavascript(),
		},
		{
			name:     "longGreatest",
			jsonData: `{"type":"longGreatest"}`,
			expected: NewLongGreatest().SetFields([]builder.PostAggregator{}),
		},
		{
			name:     "longLeast",
			jsonData: `{"type":"longLeast"}`,
			expected: NewLongLeast().SetFields([]builder.PostAggregator{}),
		},
		{
			name:     "quantileFromTDigestSketch",
			jsonData: `{"type":"quantileFromTDigestSketch"}`,
			expected: NewQuantileFromTDigestSketch(),
		},
		{
			name:     "quantilesFromTDigestSketch",
			jsonData: `{"type":"quantilesFromTDigestSketch"}`,
			expected: NewQuantilesFromTDigestSketch(),
		},
		{
			name:     "quantilesDoublesSketchToQuantile",
			jsonData: `{"type":"quantilesDoublesSketchToQuantile"}`,
			expected: NewQuantilesDoublesSketchToQuantile(),
		},
		{
			name:     "quantilesDoublesSketchToQuantiles",
			jsonData: `{"type":"quantilesDoublesSketchToQuantiles"}`,
			expected: NewQuantilesDoublesSketchToQuantiles(),
		},
		{
			name:     "quantilesDoublesSketchToHistogram",
			jsonData: `{"type":"quantilesDoublesSketchToHistogram"}`,
			expected: NewQuantilesDoublesSketchToHistogram(),
		},
		{
			name:     "quantilesDoublesSketchToRank",
			jsonData: `{"type":"quantilesDoublesSketchToRank"}`,
			expected: NewQuantilesDoublesSketchToRank(),
		},
		{
			name:     "quantilesDoublesSketchToCDF",
			jsonData: `{"type":"quantilesDoublesSketchToCDF"}`,
			expected: NewQuantilesDoublesSketchToCDF(),
		},
		{
			name:     "quantilesDoublesSketchToString",
			jsonData: `{"type":"quantilesDoublesSketchToString"}`,
			expected: NewQuantilesDoublesSketchToString(),
		},
		{
			name:     "unknown post aggregator",
			jsonData: `{"type": "blahblahType"}`,
			expected: builder.NewSpec("blahblahType"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			a, err := Load([]byte(test.jsonData))

			assert.NoErrorf(err, "loading a valid post aggregation should not fail")
			assert.Equal(test.expected, a)
		})
	}
}
