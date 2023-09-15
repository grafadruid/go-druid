package aggregation

import (
	"testing"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/filter"

	"github.com/stretchr/testify/assert"
)

func TestLoadInvalidJSON(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("not JSON value"))

	assert.Nilf(f, "aggregation should be nil")
	assert.Errorf(err, "loading should return a non-nil error")
}

func TestLoadUntypedJSON(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{}"))

	assert.Nilf(f, "aggregation should be nil")
	assert.Errorf(err, "loading should return a non-nil error")
}

func TestLoadValidAggregation(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		expected builder.Aggregator
	}{
		{
			name:     "cardinality",
			jsonData: `{"type": "cardinality"}`,
			expected: NewCardinality(),
		},
		{
			name:     "count",
			jsonData: `{"type": "count"}`,
			expected: NewCount(),
		},
		{
			name:     "doubleAny",
			jsonData: `{"type": "doubleAny"}`,
			expected: NewDoubleAny(),
		},
		{
			name:     "doubleFirst",
			jsonData: `{"type": "doubleFirst"}`,
			expected: NewDoubleFirst(),
		},
		{
			name:     "doubleLast",
			jsonData: `{"type": "doubleLast"}`,
			expected: NewDoubleLast(),
		},
		{
			name:     "doubleMax",
			jsonData: `{"type": "doubleMax"}`,
			expected: NewDoubleMax(),
		},
		{
			name:     "doubleMean",
			jsonData: `{"type": "doubleMean"}`,
			expected: NewDoubleMean(),
		},
		{
			name:     "doubleMin",
			jsonData: `{"type": "doubleMin"}`,
			expected: NewDoubleMin(),
		},
		{
			name:     "doubleSum",
			jsonData: `{"type": "doubleSum"}`,
			expected: NewDoubleSum(),
		},
		{
			name: "filtered",
			jsonData: `{
				"type": "filtered",
				"aggregator": {"type":"count"},
				"filter": {"type": "true"}
			}`,
			expected: NewFiltered().
				SetAggregator(NewCount()).
				SetFilter(filter.NewTrue()),
		},
		{
			name:     "floatAny",
			jsonData: `{"type": "floatAny"}`,
			expected: NewFloatAny(),
		},
		{
			name:     "floatFirst",
			jsonData: `{"type": "floatFirst"}`,
			expected: NewFloatFirst(),
		},
		{
			name:     "floatLast",
			jsonData: `{"type": "floatLast"}`,
			expected: NewFloatLast(),
		},
		{
			name:     "floatMax",
			jsonData: `{"type": "floatMax"}`,
			expected: NewFloatMax(),
		},
		{
			name:     "floatMin",
			jsonData: `{"type": "floatMin"}`,
			expected: NewFloatMin(),
		},
		{
			name:     "floatSum",
			jsonData: `{"type": "floatSum"}`,
			expected: NewFloatSum(),
		},
		{
			name:     "histogram",
			jsonData: `{"type": "histogram"}`,
			expected: NewHistogram(),
		},
		{
			name:     "HLLSketchBuild",
			jsonData: `{"type": "HLLSketchBuild"}`,
			expected: NewHLLSketchBuild(),
		},
		{
			name:     "HLLSketchMerge",
			jsonData: `{"type": "HLLSketchMerge"}`,
			expected: NewHLLSketchMerge(),
		},
		{
			name:     "hyperUnique",
			jsonData: `{"type": "hyperUnique"}`,
			expected: NewHyperUnique(),
		},
		{
			name:     "javascript",
			jsonData: `{"type": "javascript"}`,
			expected: NewJavascript(),
		},
		{
			name:     "longAny",
			jsonData: `{"type": "longAny"}`,
			expected: NewLongAny(),
		},
		{
			name:     "longFirst",
			jsonData: `{"type": "longFirst"}`,
			expected: NewLongFirst(),
		},
		{
			name:     "longLast",
			jsonData: `{"type": "longLast"}`,
			expected: NewLongLast(),
		},
		{
			name:     "longMax",
			jsonData: `{"type": "longMax"}`,
			expected: NewLongMax(),
		},
		{
			name:     "longMin",
			jsonData: `{"type": "longMin"}`,
			expected: NewLongMin(),
		},
		{
			name:     "longSum",
			jsonData: `{"type": "longSum"}`,
			expected: NewLongSum(),
		},
		{
			name:     "stringAny",
			jsonData: `{"type": "stringAny"}`,
			expected: NewStringAny(),
		},
		{
			name:     "stringFirstFolding",
			jsonData: `{"type": "stringFirstFolding"}`,
			expected: NewStringFirstFolding(),
		},
		{
			name:     "stringFirst",
			jsonData: `{"type": "stringFirst"}`,
			expected: NewStringFirst(),
		},
		{
			name:     "stringLastFolding",
			jsonData: `{"type": "stringLastFolding"}`,
			expected: NewStringLastFolding(),
		},
		{
			name:     "stringLast",
			jsonData: `{"type": "stringLast"}`,
			expected: NewStringLast(),
		},
		{
			name:     "tDigestSketch",
			jsonData: `{"type": "tDigestSketch"}`,
			expected: NewTDigestSketch(),
		},
		{
			name:     "quantilesDoublesSketch",
			jsonData: `{"type": "quantilesDoublesSketch"}`,
			expected: NewQuantilesDoublesSketch(),
		},
		{
			name:     "thetaSketch",
			jsonData: `{"type": "thetaSketch"}`,
			expected: NewThetaSketch(),
		},
		{
			name:     "unknown aggregator",
			jsonData: `{"type": "blahblahType"}`,
			expected: builder.NewSpec("blahblahType"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			a, err := Load([]byte(test.jsonData))

			assert.NoErrorf(err, "loading a valid aggregation should not fail")
			assert.Equal(test.expected, a)
		})
	}
}
