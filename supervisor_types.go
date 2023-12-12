package druid

import (
	"encoding/json"
	"fmt"
)

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

// BitmapFactory is a field of IndexSpec.
type BitmapFactory struct {
	Type string `json:"type"`
}

// StringEncodingStrategy type for specifying string encoding at indexing stage.
type StringEncodingStrategy struct {
	Type string `json:"type"`
	// FrontCoded fields
	BucketSize    int `json:"bucketSize,omitempty"`
	FormatVersion int `json:"formatVersion,omitempty"`
}

// IndexSpec defines segment storage format options to be used at indexing time.
// https://druid.apache.org/docs/latest/ingestion/ingestion-spec#indexspec
type IndexSpec struct {
	Bitmap                 *BitmapFactory          `json:"bitmap,omitempty"`
	DimensionCompression   string                  `json:"dimensionCompression"`
	StringEncodingStrategy *StringEncodingStrategy `json:"stringEncodingStrategy,omitempty"`
	MetricCompression      string                  `json:"metricCompression,omitempty"`
	LongEncoding           string                  `json:"longEncoding,omitempty"`
	JsonCompression        string                  `json:"jsonCompression,omitempty"`
	SegmentLoader          string                  `json:"segmentLoader,omitempty"`
}

// TuningConfig controls various tuning parameters specific to each ingestion method.
// https://druid.apache.org/docs/latest/ingestion/ingestion-spec#tuningconfig
type TuningConfig struct {
	Type                             string     `json:"type"`
	IntermediatePersistPeriod        string     `json:"intermediatePersistPeriod,omitempty"`
	MaxRowsPerSegment                int        `json:"maxRowsPerSegment,omitempty"`
	MaxRowsInMemory                  int        `json:"maxRowsInMemory,omitempty"`
	IndexSpecForIntermediatePersists *IndexSpec `json:"indexSpecForIntermediatePersists"`
}

// Metric is a Druid aggregator that is applied at ingestion time.
// https://druid.apache.org/docs/latest/ingestion/ingestion-spec#metricsspec
type Metric struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	FieldName string `json:"fieldName"`
}

// DataSchema represents the Druid dataSchema spec.
type DataSchema struct {
	DataSource      string           `json:"dataSource"`
	Parser          string           `json:"parser,omitempty"`
	TimeStampSpec   *TimestampSpec   `json:"timestampSpec,omitempty"`
	TransformSpec   *TransformSpec   `json:"transformSpec,omitempty"`
	DimensionsSpec  *DimensionsSpec  `json:"dimensionsSpec,omitempty"`
	GranularitySpec *GranularitySpec `json:"granularitySpec,omitempty"`
	MetricSpec      []Metric         `json:"metricSpec,omitempty"`
}

// FlattenSpec is responsible for flattening nested input JSON data into Druid's flat data model.
type FlattenSpec struct {
	UseFieldDiscovery bool      `json:"useFieldDiscovery,omitempty"`
	Fields            FieldList `json:"fields"`
}

// TimestampSpec is responsible for configuring the primary timestamp.
type TimestampSpec struct {
	Column string `json:"column"`
	Format string `json:"format"`
}

// FieldList is a list of Fields for ingestion FlattenSpec.
type FieldList []Field

// Field defines a single filed configuration of the FlattenSpec.
type Field struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Expr string `json:"expr"`
}

// Transform defines a single filed transformation of the TransformSpec.
type Transform struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Expr string `json:"expression"`
}

// SpatialDimension represents single spatial dimension datum.
// https://druid.apache.org/docs/latest/querying/geo/#spatial-indexing
type SpatialDimension struct {
	DimensionName string   `json:"dimName"`
	Dimensions    []string `json:"dims,omitempty"`
}

// TransformSet is a unique set of transforms applied to the input.
type TransformSet []Transform

// DimensionSet is a unique set of druid datasource dimensions(labels).
type DimensionSet []DimensionSpec

// DimensionSpec is a single dataset dimension that can be represented by a typed Dimension or a string value.
type DimensionSpec struct {
	Value any
}

// Dimension is a typed definition of a datasource dimension.
type Dimension struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// SpatialDimensionSet is a unique set of druid datasource spatial dimensions.
type SpatialDimensionSet []SpatialDimension

// DimensionExclusionsSet represents set of excluded dimensions.
type DimensionExclusionsSet []string

// DimensionsSpec is responsible for configuring Druid's dimensions. They're a
// set of columns in Druid's data model that can be used for grouping, filtering
// or applying aggregations.
// https://druid.apache.org/docs/latest/ingestion/ingestion-spec#dimensionsspec
type DimensionsSpec struct {
	Dimensions           DimensionSet           `json:"dimensions,omitempty"`
	DimensionExclusions  DimensionExclusionsSet `json:"dimensionExclusions,omitempty"`
	SpatialDimensions    SpatialDimensionSet    `json:"spatialDimensions,omitempty"`
	IncludeAllDimensions bool                   `json:"includeAllDimensions,omitempty"`
	UseSchemaDiscovery   bool                   `json:"useSchemaDiscovery,omitempty"`
}

// QueryGranularitySpec is an umbrella type for different representations of query granularity, can be string or
// QueryGranularity value.
type QueryGranularitySpec struct {
	Value any
}

// QueryGranularity is a typed representation of query granularity.
type QueryGranularity struct {
	Type string `json:"type,omitempty"`
}

// GranularitySpec allows for configuring operations such as data segment
// partitioning, truncating timestamps, time chunk segmentation or roll-up.
// https://druid.apache.org/docs/latest/ingestion/ingestion-spec#granularityspec
type GranularitySpec struct {
	Type               string                `json:"type"`
	SegmentGranularity string                `json:"segmentGranularity,omitempty"`
	QueryGranularity   *QueryGranularitySpec `json:"queryGranularity,omitempty"`
	Rollup             bool                  `json:"rollup,omitempty"`
	Intervals          []string              `json:"intervals,omitempty"`
}

// AutoScalerConfig is part of IOConfig that controls ingestion auto-scaling.
type AutoScalerConfig struct {
	EnableTaskAutoScaler                 bool    `json:"enableTaskAutoScaler"`
	LagCollectionIntervalMillis          int     `json:"lagCollectionIntervalMillis"`
	LagCollectionRangeMillis             int     `json:"lagCollectionRangeMillis"`
	ScaleOutThreshold                    int     `json:"scaleOutThreshold"`
	TriggerScaleOutFractionThreshold     float64 `json:"triggerScaleOutFractionThreshold"`
	ScaleInThreshold                     int     `json:"scaleInThreshold"`
	TriggerScaleInFractionThreshold      float64 `json:"triggerScaleInFractionThreshold"`
	ScaleActionStartDelayMillis          int     `json:"scaleActionStartDelayMillis"`
	ScaleActionPeriodMillis              int     `json:"scaleActionPeriodMillis"`
	TaskCountMax                         int     `json:"taskCountMax"`
	TaskCountMin                         int     `json:"taskCountMin"`
	ScaleInStep                          int     `json:"scaleInStep"`
	ScaleOutStep                         int     `json:"scaleOutStep"`
	MinTriggerScaleActionFrequencyMillis int     `json:"minTriggerScaleActionFrequencyMillis"`
}

// IdleConfig defines if and when stream Supervisor can become idle.
type IdleConfig struct {
	Enabled             bool  `json:"enabled"`
	InactiveAfterMillis int64 `json:"inactiveAfterMillis"`
}

// Firehose is an IOConfig firehose configuration.
type Firehose struct {
	Type string `json:"type,omitempty"`

	// EventReceiverFirehoseFactory fields
	ServiceName string `json:"serviceName,omitempty"`
	BufferSize  int    `json:"bufferSize,omitempty"`
	MaxIdleTime int64  `json:"maxIdleTime,omitempty"`

	// FixedCountFirehoseFactory / ClippedFirehoseFactory / TimedShutoffFirehoseFactory fields
	Delegate    []Firehose `json:"delegate,omitempty"`
	Count       int        `json:"count,omitempty"`
	Interval    string     `json:"interval,omitempty"`
	ShutoffTime string     `json:"shutoffTime,omitempty"`
}

// CompactionInputSpec is a specification for compaction task.
type CompactionInputSpec struct {
	Type string `json:"type"`
	// CompactionIntervalSpec fields
	Interval                 string `json:"interval,omitempty"`
	Sha256OfSortedSegmentIds string `json:"sha256OfSortedSegmentIds,omitempty"`
	// SpecificSegmentsSpec fields
	Segments []string `json:"segments,omitempty"`
}

// MetadataStorageUpdaterJobSpec is a specification of endpoint for HadoopIOConfig.
type MetadataStorageUpdaterJobSpec struct {
	Type           string         `json:"type"`
	ConnectURI     string         `json:"connectURI"`
	User           string         `json:"user"`
	Password       string         `json:"password"`
	SegmentTable   string         `json:"segmentTable"`
	CreteTable     bool           `json:"creteTable"`
	Host           string         `json:"host"`
	Port           string         `json:"port"`
	DBCPProperties map[string]any `json:"dbcp"`
}

// IOConfig influences how data is read into Druid from a source system.
// https://druid.apache.org/docs/latest/ingestion/ingestion-spec/#ioconfig
type IOConfig struct {
	Type string `json:"type,omitempty"`

	// IndexIOConfig / RealtimeIOConfig shared field
	Firehose *Firehose `json:"firehose,omitempty"`
	// IndexIOConfig field
	InputSource      *InputSource `json:"inputSource,omitempty"`
	AppendToExisting bool         `json:"appendToExisting,omitempty"`
	// IndexIOConfig / CompactionIOConfig shared field
	DropExisting bool `json:"dropExisting,omitempty"`

	// CompactionIOConfig / HadoopIOConfig fields
	InputSpec map[string]any `json:"inputSpec,omitempty"`

	// CompactionIOConfig field
	AllowNonAlignedInterval bool `json:"allowNonAlignedInterval,omitempty"`

	// HadoopIOConfig fields
	MetadataUpdateSpec *MetadataStorageUpdaterJobSpec `json:"metadataUpdateSpec,omitempty"`
	SegmentOutputPath  string                         `json:"segmentOutputPath,omitempty"`

	// KafkaIndexTaskIOConfig / KinesisIndexTaskIOConfig fields
	Topic                     string              `json:"topic,omitempty"`
	ConsumerProperties        *ConsumerProperties `json:"consumerProperties,omitempty"`
	TaskDuration              string              `json:"taskDuration,omitempty"`
	Replicas                  int                 `json:"replicas,omitempty"`
	TaskCount                 int                 `json:"taskCount,omitempty"`
	UseEarliestOffset         bool                `json:"useEarliestOffset"`
	AutoScalerConfig          *AutoScalerConfig   `json:"autoScalerConfig,omitempty"`
	TaskGroupID               int                 `json:"taskGroupID,omitempty"`
	BaseSequenceName          string              `json:"baseSequenceName,omitempty"`
	CompletionTimeout         string              `json:"completionTimeout,omitempty"`
	PollTimeout               int                 `json:"pollTimeout,omitempty"`
	StartDelay                string              `json:"startDelay,omitempty"`
	Period                    string              `json:"period,omitempty"`
	Stream                    string              `json:"stream,omitempty"`
	UseEarliestSequenceNumber bool                `json:"useEarliestSequenceNumber,omitempty"`

	// common fields
	FlattenSpec *FlattenSpec `json:"flattenSpec,omitempty"`
	InputFormat *InputFormat `json:"inputFormat,omitempty"`
	IdleConfig  *IdleConfig  `json:"idleConfig,omitempty"`
}

// ConsumerProperties is a set of properties that is passed to a specific
// consumer, i.e. Kafka consumer.
type ConsumerProperties struct {
	BootstrapServers string `json:"bootstrap.servers,omitempty"`
}

// InputFormat specifies kafka messages format type and describes any conversions applied to
// the input data while parsing.
// Type can take values 'json', 'protobuf' or 'kafka'.
type InputFormat struct {
	Type string `json:"type"`

	// FlatTextInputFormat / DelimitedInputFormat fields
	Delimiter         string   `json:"delimiter,omitempty"`
	ListDelimiter     string   `json:"listDelimiter,omitempty"`
	FindColumnsHeader string   `json:"findColumnsHeader,omitempty"`
	SkipHeaderRows    int      `json:"skipHeaderRows,omitempty"`
	Columns           []string `json:"columns,omitempty"`

	// JsonInputFormat fields
	FlattenSpec *FlattenSpec    `json:"flattenSpec,omitempty"`
	FeatureSpec map[string]bool `json:"featureSpec,omitempty"`

	// Common CsvInputFormat / JsonInputFormat fields
	KeepNullColumns        bool `json:"keepNullColumns,omitempty"`
	AssumeNewlineDelimited bool `json:"assumeNewlineDelimited,omitempty"`
	UseJsonNodeReader      bool `json:"useJsonNodeReader,omitempty"`
}

// HttpInputSourceConfig is a field of HttpInputSource specification.
type HttpInputSourceConfig struct {
	AllowedProtocols []string `json:" allowedProtocols,omitempty"`
}

// ConnectorConfig is connection configuration for Database.
type ConnectorConfig struct {
	ConnectURI string `json:"connectURI"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

// Database configuration for InputSource "sql".
type Database struct {
	Type            string           `json:"type"`
	ConnectorConfig *ConnectorConfig `json:"connectorConfig"`
}

// InputSource  is a specification of the storage system where input data is stored.
type InputSource struct {
	Type string `json:"type"`

	// LocalInputSource fields
	BaseDir string   `json:"baseDir,omitempty"`
	Filter  string   `json:"filter,omitempty"`
	Files   []string `json:"files,omitempty"`

	// HttpInputSource fields
	URIs                       []string               `json:"uris,omitempty"`
	HttpAuthenticationUsername string                 `json:"httpAuthenticationUsername,omitempty"`
	HttpAuthenticationPassword string                 `json:"httpAuthenticationPassword,omitempty"`
	HttpSourceConfig           *HttpInputSourceConfig `json:"config,omitempty"`

	// InlineInputSource fields
	Data string `json:"data,omitempty"`

	// CombiningInputSource fields
	Delegates []InputSource `json:"delegates,omitempty"`

	// SqlInputSource
	SQLs     []string  `json:"sqls,omitempty"`
	Database *Database `json:"database,omitempty"`
}

// TransformSpec is responsible for transforming druid input data
// after it was read from kafka and after flattenSpec was applied.
// https://druid.apache.org/docs/latest/ingestion/ingestion-spec#transformspec
type TransformSpec struct {
	Transforms TransformSet `json:"transforms"`
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

func (g *QueryGranularitySpec) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		g.Value = str
		return nil
	}

	var qg QueryGranularity
	if err := json.Unmarshal(b, &qg); err == nil {
		g.Value = qg
		return nil
	}
	return fmt.Errorf("unsupported query granularity: %s", b)
}

func (g *QueryGranularitySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(&g.Value)
}

func (g *DimensionSpec) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		g.Value = str
		return nil
	}

	var qg Dimension
	if err := json.Unmarshal(b, &qg); err == nil {
		g.Value = qg
		return nil
	}
	return fmt.Errorf("unsupported dimension value: %s", b)
}

func (g *DimensionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(&g.Value)
}
