package druid

// InputIngestionSpec is the root-level type defining an ingestion spec used
// by Apache Druid.
type InputIngestionSpec struct {
	Type         string        `json:"type"`
	DataSchema   *DataSchema   `json:"dataSchema,omitempty"`
	IOConfig     *IOConfig     `json:"ioConfig,omitempty"`
	TuningConfig *TuningConfig `json:"tuningConfig,omitempty"`
}

// IngestionSpecData is the core supervisor specification data returned by druid supervisor APIs.
// It is a part of OutputIngestionSpec.
type IngestionSpecData struct {
	DataSchema   *DataSchema   `json:"dataSchema,omitempty"`
	IOConfig     *IOConfig     `json:"ioConfig,omitempty"`
	TuningConfig *TuningConfig `json:"tuningConfig,omitempty"`
}

// OutputIngestionSpec is full supervisor specification format returned by druid supervisor APIs.
type OutputIngestionSpec struct {
	Type      string             `json:"type"`
	Context   string             `json:"context"`
	Suspended bool               `json:"suspended"`
	Spec      *IngestionSpecData `json:"spec"`
}

// SupervisorAuditHistory is audit data for supervisor reurned by supervisor audit history APIs.
type SupervisorAuditHistory struct {
	Spec    OutputIngestionSpec `json:"spec"`
	Version string              `json:"version"`
}

// SupervisorStatusPayload is an object representing the status of supervisor.
type SupervisorStatusPayload struct {
	Datasource      string `json:"dataSource"`
	Stream          string `json:"stream"`
	State           string `json:"state"`
	Partitions      int    `json:"partitions"`
	Replicas        int    `json:"replicas"`
	DurationSeconds int    `json:"durationSeconds"`
	Suspended       bool   `json:"suspended"`
}

// SupervisorStatus is a response object containing status of a supervisor alongside
// with the response metadata.
type SupervisorStatus struct {
	SupervisorId   string                   `json:"id"`
	GenerationTime string                   `json:"generationTime"`
	Payload        *SupervisorStatusPayload `json:"payload"`
}

// SupervisorState is a short form of supervisor state returned by druid APIs.
type SupervisorState struct {
	ID            string `json:"id"`
	State         string `json:"state"`
	DetailedState string `json:"detailedState"`
	Healthy       bool   `json:"healthy"`
	Suspended     bool   `json:"suspended"`
}

type SupervisorStateWithSpec struct {
	SupervisorState
	Spec *InputIngestionSpec `json:"spec"`
}

// defaultKafkaIngestionSpec returns a default InputIngestionSpec with basic ingestion
// specification fields initialized.
func defaultKafkaIngestionSpec() *InputIngestionSpec {
	spec := &InputIngestionSpec{
		Type: "kafka",
		DataSchema: &DataSchema{
			DataSource: "test",
			TimeStampSpec: &TimestampSpec{
				Column: "ts",
				Format: "auto",
			},
			TransformSpec: &TransformSpec{
				Transforms: []Transform{},
			},
			DimensionsSpec: &DimensionsSpec{
				Dimensions: DimensionSet{},
			},
			GranularitySpec: &GranularitySpec{
				Type:               "uniform",
				SegmentGranularity: "DAY",
				QueryGranularity:   &QueryGranularitySpec{"none"},
				Rollup:             false,
			},
		},
		IOConfig: &IOConfig{
			Type:  "",
			Topic: "",
			InputFormat: &InputFormat{
				Type: "json",
			},
			TaskDuration: "PT1H",
			ConsumerProperties: &ConsumerProperties{
				BootstrapServers: "",
			},
			UseEarliestOffset: false,
			FlattenSpec: &FlattenSpec{
				Fields: FieldList{},
			},
		},
	}
	return spec
}

// NewIngestionSpec returns a default InputIngestionSpec and applies any
// options passed to it.
func NewIngestionSpec(options ...IngestionSpecOptions) *InputIngestionSpec {
	spec := defaultKafkaIngestionSpec()
	for _, fn := range options {
		fn(spec)
	}
	return spec
}

// IngestionSpecOptions allows for configuring a InputIngestionSpec.
type IngestionSpecOptions func(*InputIngestionSpec)

// SetType sets the type of the supervisor (IOConfig).
func SetType(stype string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if stype != "" {
			spec.Type = stype
		}
	}
}

// SetIOConfigType sets the type of the supervisor IOConfig.
func SetIOConfigType(ioctype string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if ioctype != "" {
			spec.IOConfig.Type = ioctype
		}
	}
}

// SetTopic sets the Kafka topic to consume data from.
func SetTopic(topic string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if topic != "" {
			spec.IOConfig.Topic = topic
		}
	}
}

// SetDataSource sets the name of the dataSource used in Druid.
func SetDataSource(ds string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if ds != "" {
			spec.DataSchema.DataSource = ds
		}
	}
}

// SetInputFormat sets the input format type, i.e. json, protobuf etc.
func SetInputFormat(format string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if format != "" {
			spec.IOConfig.InputFormat.Type = format
		}
	}
}

// SetBrokers sets the addresses of Kafka brokers. in the list form: 'kafka01:9092,
// kafka02:9092,kafka03:9092' or as a cluster DNS: kafka.default.svc.cluster.local:9092”.
func SetBrokers(brokers string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if brokers != "" {
			spec.IOConfig.ConsumerProperties.BootstrapServers = brokers
		}
	}
}

// SetTaskDuration sets the upper limit for druid ingestion task.
func SetTaskDuration(duration string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if duration != "" {
			spec.IOConfig.TaskDuration = duration
		}
	}
}

// SetDimensions sets druid datasource dimensions.
func SetDimensions(dimensions DimensionSet) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		spec.DataSchema.DimensionsSpec.Dimensions = dimensions
	}
}

// SetDimensionsAutodiscovery sets druid autodiscovery for datasource dimensions.
func SetDimensionsAutodiscovery(discover bool) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		spec.DataSchema.DimensionsSpec.UseSchemaDiscovery = discover
	}
}

// SetUseEarliestOffset configures kafka druid ingestion supervisor to start reading
// from the earliest or latest offsets in Kafka.
func SetUseEarliestOffset(useEarliestOffset bool) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		spec.IOConfig.UseEarliestOffset = useEarliestOffset
	}
}

// SetTimestampColumn sets timestamp column for the druid datasource.
func SetTimestampColumn(column string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		if column != "" {
			spec.DataSchema.TimeStampSpec = &TimestampSpec{
				Column: column,
				Format: "auto",
			}
		}
	}
}

// SetGranularitySpec sets granularity spec settings that are applied at druid ingestion partitioning stage.
func SetGranularitySpec(segmentGranularity string, queryGranularity *QueryGranularitySpec, rollup bool) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		spec.DataSchema.GranularitySpec = &GranularitySpec{
			Type:               "uniform",
			SegmentGranularity: segmentGranularity,
			QueryGranularity:   queryGranularity,
			Rollup:             rollup,
		}
	}
}

// SetSQLInputSource configures sql input source.
func SetSQLInputSource(dbType, connectURI, user, password string, sqls []string) IngestionSpecOptions {
	return func(spec *InputIngestionSpec) {
		spec.IOConfig.InputSource = &InputSource{
			Type: "sql",
			SQLs: sqls,
			Database: &Database{
				Type: dbType,
				ConnectorConfig: &ConnectorConfig{
					ConnectURI: connectURI,
					User:       user,
					Password:   password,
				},
			},
		}
	}
}
