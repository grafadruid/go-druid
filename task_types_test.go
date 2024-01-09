package druid

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskIngestionSpec(t *testing.T) {
	var testData = []struct {
		name     string
		options  []TaskIngestionSpecOptions
		expected *TaskIngestionSpec
	}{
		{
			name: "druid reindex ingestion task",
			options: []TaskIngestionSpecOptions{
				SetTaskDataSource("telemetry-test"),
				SetTaskDataDimensions(DimensionSet{{"id"}, {"kind"}, {"test"}}),
				SetTaskDruidInputSource(
					"telemetry-test",
					time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC),
				),
				SetTaskSchemaDiscovery(false),
				SetTaskTimestampColumn("__time"),
				SetTaskGranularitySpec("DAY", &QueryGranularitySpec{"none"}, true),
			},
			expected: func() *TaskIngestionSpec {
				out := NewTaskIngestionSpec()
				out.Spec.DataSchema.DataSource = "telemetry-test"
				out.Spec.DataSchema.DimensionsSpec = &DimensionsSpec{
					Dimensions: DimensionSet{{"id"}, {"kind"}, {"test"}},
				}
				out.Spec.IOConfig.InputSource.Type = "druid"
				out.Spec.IOConfig.InputSource.Datasource = "telemetry-test"
				out.Spec.IOConfig.InputSource.Interval = "2023-12-12T00:00:00/2023-12-24T00:00:00"
				out.Spec.DataSchema.TimeStampSpec.Column = "__time"
				out.Spec.DataSchema.TimeStampSpec.Format = "auto"
				out.Spec.DataSchema.GranularitySpec.SegmentGranularity = "DAY"
				out.Spec.DataSchema.GranularitySpec.QueryGranularity = &QueryGranularitySpec{"none"}
				out.Spec.DataSchema.GranularitySpec.Rollup = true
				return out
			}(),
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			actual := NewTaskIngestionSpec(
				test.options...,
			)
			assert.Equal(t, test.expected, actual)
		})
	}
}
