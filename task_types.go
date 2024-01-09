package druid

import (
	"fmt"
	"time"
)

// TaskStatusResponse is a response object containing status of a task.
type TaskStatusResponse struct {
	Task   string     `json:"task"`
	Status TaskStatus `json:"status"`
}

// TaskLocation holds location of the task execution.
type TaskLocation struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	TlsPort int    `json:"tlsPort"`
}

// TaskStatus is an object representing status of a druid task.
type TaskStatus struct {
	ID                 string        `json:"id"`
	Type               string        `json:"type"`
	CreatedTime        string        `json:"createdTime"`
	QueueInsertionTime string        `json:"queueInsertionTime"`
	StatusCode         string        `json:"statusCode"`
	Status             string        `json:"status"`
	RunnerStatusCode   string        `json:"runnerStatusCode"`
	Duration           int           `json:"duration"`
	GroupId            string        `json:"groupId"`
	Location           *TaskLocation `json:"location|omitempty"`
	Datasource         string        `json:"datasource"`
	ErrorMessage       string        `json:"errorMessage"`
}

// TaskIngestionSpec is a specification for a druid task execution.
type TaskIngestionSpec struct {
	Type string             `json:"type"`
	Spec *IngestionSpecData `json:"spec"`
}

// defaultTaskIngestionSpec returns a default TaskIngestionSpec with basic ingestion
// specification fields initialized.
func defaultTaskIngestionSpec() *TaskIngestionSpec {
	spec := &TaskIngestionSpec{
		Type: "index_parallel",
		Spec: &IngestionSpecData{
			DataSchema: &DataSchema{
				DataSource: "some_datasource",
				GranularitySpec: &GranularitySpec{
					Type:               "uniform",
					SegmentGranularity: "DAY",
					QueryGranularity:   &QueryGranularitySpec{"none"},
				},
				DimensionsSpec: &DimensionsSpec{
					UseSchemaDiscovery: true,
					Dimensions:         DimensionSet{},
				},
				TransformSpec: &TransformSpec{
					Transforms: []Transform{},
				},
				TimeStampSpec: &TimestampSpec{},
			},
			IOConfig: &IOConfig{
				Type:        "index_parallel",
				InputSource: &InputSource{},
				InputFormat: &InputFormat{},
			},
			TuningConfig: &TuningConfig{
				Type: "index_parallel",
			},
		},
	}
	return spec
}

// TaskIngestionSpecOptions allows for configuring a TaskIngestionSpec.
type TaskIngestionSpecOptions func(*TaskIngestionSpec)

// SetTaskType sets the type of the task IOConfig.
func SetTaskType(stype string) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		if stype != "" {
			spec.Type = stype
		}
	}
}

// SetTaskTimestampColumn sets the type of the task IOConfig.
func SetTaskTimestampColumn(column string) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		if column != "" {
			spec.Spec.DataSchema.TimeStampSpec = &TimestampSpec{
				Column: column,
				Format: "auto",
			}
		}
	}
}

// SetTaskDataSource sets the destination datasource of the task IOConfig.
func SetTaskDataSource(datasource string) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		if datasource != "" {
			spec.Spec.DataSchema.DataSource = datasource
		}
	}
}

// SetTaskTuningConfig sets the tuning configuration the task IOConfig.
func SetTaskTuningConfig(typ string, maxRowsInMemory, maxRowsPerSegment int) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		if typ != "" {
			spec.Spec.TuningConfig.Type = typ
			spec.Spec.TuningConfig.MaxRowsInMemory = maxRowsInMemory
			spec.Spec.TuningConfig.MaxRowsPerSegment = maxRowsPerSegment
		}
	}
}

// SetTaskDataDimensions sets druid datasource dimensions.
func SetTaskDataDimensions(dimensions DimensionSet) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		spec.Spec.DataSchema.DimensionsSpec.Dimensions = dimensions
	}
}

// SetTaskSQLInputSource configures sql input source for the task based ingestion.
func SetTaskSQLInputSource(typ, connectURI, user, password string, sqls []string) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		spec.Spec.IOConfig.InputSource = &InputSource{
			Type: "sql",
			SQLs: sqls,
			Database: &Database{
				Type: typ,
				ConnectorConfig: &ConnectorConfig{
					ConnectURI: connectURI,
					User:       user,
					Password:   password,
				},
			},
		}
	}
}

// SetTaskIOConfigType sets the type of the task IOConfig.
func SetTaskIOConfigType(typ string) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		if typ != "" {
			spec.Spec.IOConfig.Type = typ
		}
	}
}

// SetTaskInputFormat configures input format for the task based ingestion.
func SetTaskInputFormat(typ string, findColumnsHeader string, columns []string) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		spec.Spec.IOConfig.InputFormat.Type = typ
		spec.Spec.IOConfig.InputFormat.FindColumnsFromHeader = findColumnsHeader
		spec.Spec.IOConfig.InputFormat.Columns = columns
	}
}

// SetTaskInlineInputData configures inline data for the task based ingestion.
func SetTaskInlineInputData(data string) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		spec.Spec.IOConfig.InputSource.Type = "inline"
		spec.Spec.IOConfig.InputSource.Data = data
	}
}

// SetTaskDruidInputSource configures druid reindex input source for the task based ingestion.
func SetTaskDruidInputSource(datasource string, startTime time.Time, endTime time.Time) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		spec.Spec.IOConfig.InputSource = &InputSource{
			Type:       "druid",
			Datasource: datasource,
			Interval: fmt.Sprintf(
				"%s/%s",
				startTime.Format("2006-01-02T15:04:05"),
				endTime.Format("2006-01-02T15:04:05"),
			),
		}
	}
}

// SetTaskSchemaDiscovery sets auto discovery of dimensions.
func SetTaskSchemaDiscovery(discovery bool) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		spec.Spec.DataSchema.DimensionsSpec.UseSchemaDiscovery = discovery
	}
}

// SetTaskGranularitySpec sets granularity spec settings that are applied at druid ingestion partitioning stage.
func SetTaskGranularitySpec(segmentGranularity string, queryGranularity *QueryGranularitySpec, rollup bool) TaskIngestionSpecOptions {
	return func(spec *TaskIngestionSpec) {
		spec.Spec.DataSchema.GranularitySpec = &GranularitySpec{
			Type:               "uniform",
			SegmentGranularity: segmentGranularity,
			QueryGranularity:   queryGranularity,
			Rollup:             rollup,
		}
	}
}

// NewTaskIngestionSpec returns a default TaskIngestionSpec and applies any options passed to it.
func NewTaskIngestionSpec(options ...TaskIngestionSpecOptions) *TaskIngestionSpec {
	spec := defaultTaskIngestionSpec()
	for _, fn := range options {
		fn(spec)
	}
	return spec
}
