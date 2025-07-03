package postaggregation

import (
	"github.com/grafadruid/go-druid/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadAggregation(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected builder.Aggregator
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "empty json",
			jsonData: "{}",
			wantErr:  assert.Error,
		},
		{
			name:     "invalid json",
			jsonData: "not JSON value",
			wantErr:  assert.Error,
		},
		{
			name:     "arithmetic",
			jsonData: `{"type":"arithmetic"}`,
			expected: NewArithmetic().SetFields([]builder.PostAggregator{}),
			wantErr:  assert.NoError,
		},
		{
			name:     "constant",
			jsonData: `{"type":"constant"}`,
			expected: NewConstant(),
			wantErr:  assert.NoError,
		},
		{
			name:     "doubleGreatest",
			jsonData: `{"type":"doubleGreatest"}`,
			expected: NewDoubleGreatest().SetFields([]builder.PostAggregator{}),
			wantErr:  assert.NoError,
		},
		{
			name:     "doubleLeast",
			jsonData: `{"type":"doubleLeast"}`,
			expected: NewDoubleLeast().SetFields([]builder.PostAggregator{}),
			wantErr:  assert.NoError,
		},
		{
			name:     "expression",
			jsonData: `{"type":"expression"}`,
			expected: NewExpression(),
			wantErr:  assert.NoError,
		},
		{
			name:     "fieldAccess",
			jsonData: `{"type":"fieldAccess"}`,
			expected: NewFieldAccess(),
			wantErr:  assert.NoError,
		},
		{
			name:     "finalizingFieldAccess",
			jsonData: `{"type":"finalizingFieldAccess"}`,
			expected: NewFinalizingFieldAccess(),
			wantErr:  assert.NoError,
		},
		{
			name:     "hyperUniqueFinalizing",
			jsonData: `{"type":"hyperUniqueFinalizing"}`,
			expected: NewHyperUniqueFinalizing(),
			wantErr:  assert.NoError,
		},
		{
			name:     "javascript",
			jsonData: `{"type":"javascript"}`,
			expected: NewJavascript(),
			wantErr:  assert.NoError,
		},
		{
			name:     "longGreatest",
			jsonData: `{"type":"longGreatest"}`,
			expected: NewLongGreatest().SetFields([]builder.PostAggregator{}),
			wantErr:  assert.NoError,
		},
		{
			name:     "longLeast",
			jsonData: `{"type":"longLeast"}`,
			expected: NewLongLeast().SetFields([]builder.PostAggregator{}),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantileFromTDigestSketch",
			jsonData: `{"type":"quantileFromTDigestSketch"}`,
			expected: NewQuantileFromTDigestSketch(),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantilesFromTDigestSketch",
			jsonData: `{"type":"quantilesFromTDigestSketch"}`,
			expected: NewQuantilesFromTDigestSketch(),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantilesDoublesSketchToQuantile",
			jsonData: `{"type":"quantilesDoublesSketchToQuantile"}`,
			expected: NewQuantilesDoublesSketchToQuantile(),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantilesDoublesSketchToQuantiles",
			jsonData: `{"type":"quantilesDoublesSketchToQuantiles"}`,
			expected: NewQuantilesDoublesSketchToQuantiles(),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantilesDoublesSketchToHistogram",
			jsonData: `{"type":"quantilesDoublesSketchToHistogram"}`,
			expected: NewQuantilesDoublesSketchToHistogram(),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantilesDoublesSketchToRank",
			jsonData: `{"type":"quantilesDoublesSketchToRank"}`,
			expected: NewQuantilesDoublesSketchToRank(),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantilesDoublesSketchToCDF",
			jsonData: `{"type":"quantilesDoublesSketchToCDF"}`,
			expected: NewQuantilesDoublesSketchToCDF(),
			wantErr:  assert.NoError,
		},
		{
			name:     "quantilesDoublesSketchToString",
			jsonData: `{"type":"quantilesDoublesSketchToString"}`,
			expected: NewQuantilesDoublesSketchToString(),
			wantErr:  assert.NoError,
		},
		{
			name:     "theta sketch estimate",
			jsonData: `{"type": "thetaSketchEstimate"}`,
			expected: NewThetaSketchEstimate(),
			wantErr:  assert.NoError,
		},
		{
			name:     "unknown post aggregator",
			jsonData: `{"type": "blahblahType"}`,
			expected: builder.NewSpec("blahblahType"),
			wantErr:  assert.NoError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Load([]byte(test.jsonData))
			test.wantErr(t, err)
			assert.Equal(t, test.expected, got)
		})
	}
}
