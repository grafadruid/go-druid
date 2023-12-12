package druid

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKafkaIngestionSpec(t *testing.T) {
	var testData = []struct {
		name     string
		options  []IngestionSpecOptions
		expected *InputIngestionSpec
	}{
		{
			name: "set datasource, topic, brokers and duration",
			options: []IngestionSpecOptions{
				SetTopic("test_topic"),
				SetDataSource("test_source"),
				SetTaskDuration("PT20S"),
				SetBrokers("test_brokers"),
			},
			expected: func() *InputIngestionSpec {
				out := defaultKafkaIngestionSpec()
				out.IOConfig.Topic = "test_topic"
				out.IOConfig.ConsumerProperties.BootstrapServers = "test_brokers"
				out.DataSchema.DataSource = "test_source"
				out.IOConfig.TaskDuration = "PT20S"
				return out
			}(),
		},
		{
			name: "set labels",
			options: []IngestionSpecOptions{
				SetDimensions(DimensionSet{
					{"ts"},
					{"user_name"},
					{"payload"},
				}),
			},
			expected: func() *InputIngestionSpec {
				out := defaultKafkaIngestionSpec()
				out.DataSchema.DimensionsSpec.Dimensions = DimensionSet{{"ts"}, {"user_name"}, {"payload"}}
				return out
			}(),
		},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			actual := NewIngestionSpec(
				test.options...,
			)
			assert.Equal(t, test.expected, actual)
		})
	}
}

var jsonBasic = `{
    "type": "kafka",
    "dataSchema": {
        "dataSource": "test_datasource",
        "timestampSpec": {
            "column": "ts",
            "format": "auto"
        },
        "transformSpec": {
            "transforms": []
        },
        "dimensionsSpec": {
            "dimensions": [
                "ts",
                "user_name",
                "payload"
            ]
        },
        "granularitySpec": {
			"type": "uniform",
            "segmentGranularity": "DAY",
            "queryGranularity":  "none"
        }
    },
    "ioConfig": {
        "topic": "test_topic",
        "consumerProperties": {
            "bootstrap.servers": "test_brokers"
        },
        "taskDuration": "PT1H",
        "useEarliestOffset": false,
        "flattenSpec": {
            "fields": []
        },
        "inputFormat": {
            "type": "json"
        }
    }
}`

func TestKafkaIngestionSpec_MarshalJSON(t *testing.T) {
	spec := NewIngestionSpec(
		SetDataSource("test_datasource"),
		SetTopic("test_topic"),
		SetBrokers("test_brokers"),
		SetDimensions(DimensionSet{
			{"ts"},
			{"user_name"},
			{"payload"},
		}),
	)
	actual, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("unexpected error while marshalling: %v", err)
	}
	expected := []byte(jsonBasic)
	require.JSONEq(t, string(expected), string(actual), fmt.Sprintf("expected: %s\nactual: %s", string(expected), string(actual)))

	var checkSpec *InputIngestionSpec
	err = json.Unmarshal(actual, &checkSpec)
	if err != nil {
		t.Fatalf("unexpected error while unmarshalling: %v", err)
	}
	require.Equal(t, spec, checkSpec)
}

var jsonWithTypedDimensions = `{
    "type": "kafka",
    "dataSchema": {
        "dataSource": "test_datasource",
        "timestampSpec": {
            "column": "ts",
            "format": "auto"
        },
        "transformSpec": {
            "transforms": []
        },
        "dimensionsSpec": {
            "dimensions": [
                {
                    "type": "string",
                    "name": "ts"
                },
                {
                    "type": "json",
                    "name": "payload"
                }
            ]
        },
        "granularitySpec": {
			"type": "uniform",
            "segmentGranularity": "DAY",
            "queryGranularity": "none"
        }
    },
    "ioConfig": {
        "topic": "test_topic",
        "consumerProperties": {
            "bootstrap.servers": "test_brokers"
        },
        "taskDuration": "PT1H",
        "useEarliestOffset": false,
        "flattenSpec": {
            "fields": []
        },
        "inputFormat": {
            "type": "json"
        }
    }
}`

func TestIngestionSpecWithTypedDimensions_MarshalJSON(t *testing.T) {
	spec := NewIngestionSpec(
		SetDataSource("test_datasource"),
		SetTopic("test_topic"),
		SetBrokers("test_brokers"),
		SetDimensions(DimensionSet{
			{Dimension{Type: "string", Name: "ts"}},
			{Dimension{Type: "json", Name: "payload"}},
		}),
	)
	actual, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("unexpected error while marshalling: %v", err)
	}
	expected := []byte(jsonWithTypedDimensions)
	require.JSONEq(t, string(expected), string(actual), fmt.Sprintf("expected: %s\nactual: %s", string(expected), string(actual)))
}

var jsonWithSqlInputSource = `{
    "type": "index_parallel",
    "dataSchema": {
        "dataSource": "test_datasource",
        "timestampSpec": {
            "column": "ts",
            "format": "auto"
        },
        "transformSpec": {
            "transforms": []
        },
        "dimensionsSpec": {
            "dimensions": [
                "ts",
                "user_name",
                "payload"
            ]
        },
        "granularitySpec": {
			"type": "uniform",
			"segmentGranularity": "DAY",
            "queryGranularity": "none"
        }
    },
    "ioConfig": {
        "type": "index_parallel",
        "inputSource": {
            "type": "sql",
            "sqls": [
                "SELECT * FROM table1 WHERE timestamp BETWEEN '2013-01-01 00:00:00' AND '2013-01-01 11:59:59'",
                "SELECT * FROM table2 WHERE timestamp BETWEEN '2013-01-01 00:00:00' AND '2013-01-01 11:59:59'"
            ],
            "database": {
                "type": "mysql",
                "connectorConfig": {
                    "connectURI": "jdbc:mysql://host:port/schema",
                    "user": "username",
                    "password": "password"
                }
            }
        },
        "consumerProperties": {},
        "taskDuration": "PT1H",
        "useEarliestOffset": false,
        "flattenSpec": {
            "fields": []
        },
        "inputFormat": {
            "type": "json"
        }
    }
}`

func TestIngestionSpecWithSqlInputSource_MarshalJSON(t *testing.T) {
	spec := NewIngestionSpec(
		SetType("index_parallel"),
		SetIOConfigType("index_parallel"),
		SetDataSource("test_datasource"),
		SetDimensions(DimensionSet{{"ts"}, {"user_name"}, {"payload"}}),
		SetSQLInputSource("mysql",
			"jdbc:mysql://host:port/schema",
			"username",
			"password",
			[]string{
				"SELECT * FROM table1 WHERE timestamp BETWEEN '2013-01-01 00:00:00' AND '2013-01-01 11:59:59'",
				"SELECT * FROM table2 WHERE timestamp BETWEEN '2013-01-01 00:00:00' AND '2013-01-01 11:59:59'",
			}),
	)
	actual, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("unexpected error while marshalling: %v", err)
	}
	expected := []byte(jsonWithSqlInputSource)

	require.JSONEq(t, string(expected), string(actual), fmt.Sprintf("expected: %s\nactual: %s", string(expected), string(actual)))

	var checkSpec *InputIngestionSpec
	err = json.Unmarshal(actual, &checkSpec)
	if err != nil {
		t.Fatalf("unexpected error while unmarshalling: %v", err)
	}
	require.Equal(t, spec, checkSpec)
}
