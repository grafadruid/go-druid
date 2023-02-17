package main

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/aggregation"
	datasource2 "github.com/grafadruid/go-druid/builder/datasource"
	"github.com/grafadruid/go-druid/builder/dimension"
	"github.com/grafadruid/go-druid/builder/extractionfn"
	"github.com/grafadruid/go-druid/builder/filter"
	"github.com/grafadruid/go-druid/builder/granularity"
	"github.com/grafadruid/go-druid/builder/intervals"
	"github.com/grafadruid/go-druid/builder/query"
	"github.com/grafadruid/go-druid/builder/topnmetric"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSubQueryScan(t *testing.T) {
	expected := `{
  "batchSize": 20480,
  "columns": [
    "__time",
    "channel",
    "cityName",
    "comment",
    "count",
    "countryIsoCode",
    "diffUrl",
    "flags",
    "isAnonymous",
    "isMinor",
    "isNew",
    "isRobot",
    "isUnpatrolled",
    "metroCode",
    "namespace",
    "page",
    "regionIsoCode",
    "regionName",
    "sum_added",
    "sum_commentLength",
    "sum_deleted",
    "sum_delta",
    "sum_deltaBucket",
    "user"
  ],
  "dataSource": {
    "type": "query",
    "query": {
      "queryType": "scan",
      "dataSource": {
        "type": "table",
        "name": "A"
      },
      "columns": [
        "AT"
      ],
      "intervals": {
        "type": "intervals",
        "intervals": [
          "1980-06-12T22:30:00Z/2020-01-26T23:00:00Z"
        ]
      }
    }
  },
  "filter": {
    "dimension": "countryName",
    "extractionFn": {
      "locale": "",
      "type": "lower"
    },
    "type": "selector",
    "value": "france"
  },
  "intervals": {
    "type": "intervals",
    "intervals": [
      "1980-06-12T22:30:00Z/2020-01-26T23:00:00Z"
    ]
  },
  "limit": 10,
  "order": "DESCENDING",
  "queryType": "scan"
}`
	location, _ := time.LoadLocation("UTC")

	start, _ := time.ParseInLocation(time.RFC3339Nano,
		"1980-06-12T22:30:00.000Z",
		location)
	end, _ := time.ParseInLocation(time.RFC3339Nano,
		"2020-01-26T23:00:00.000Z",
		location)
	i := intervals.NewInterval()
	i.SetInterval(start, end)
	interval := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})

	selectorExtractionFn := extractionfn.NewLower().SetLocale("")
	a := filter.NewSelector().SetDimension("countryName").SetValue("france").
		SetExtractionFn(selectorExtractionFn)

	datasourceSubQuery := query.NewScan().
		SetDataSource(datasource2.NewTable().SetName("A")).SetColumns([]string{"AT"}).
		SetIntervals(interval)
	datasource := datasource2.NewQuery().SetQuery(datasourceSubQuery)
	scan := query.NewScan()
	queryTest := scan.SetOrder(query.Descending).SetLimit(10).SetIntervals(interval).SetFilter(a).
		SetColumns([]string{"__time", "channel", "cityName", "comment", "count", "countryIsoCode",
			"diffUrl", "flags", "isAnonymous", "isMinor", "isNew", "isRobot", "isUnpatrolled",
			"metroCode", "namespace", "page", "regionIsoCode", "regionName", "sum_added",
			"sum_commentLength", "sum_deleted", "sum_delta", "sum_deltaBucket", "user"}).
		SetBatchSize(20480).
		SetDataSource(datasource)

	queryTestJSON, err := json.Marshal(queryTest)
	assert.Nil(t, err)

	assert.JSONEq(t, expected, string(queryTestJSON))
}

func TestSubQueryTimeseries(t *testing.T) {
	expected := `{
	"dataSource": {
		"query": {
			"aggregations": [{
				"fieldName": "count",
				"name": "count",
				"type": "longSum"
			}],
			"dataSource": {
				"name": "dc_94b4f5fdfde940979b79c50539d8322a_b42fde98efed4e638a0016b34b3c10cf_dataset_pre",
				"type": "table"
			},
			"dimension": {
				"dimension": "string_value",
				"type": "default"
			},
			"filter": {
				"fields": [{
						"dimension": "_split_name_",
						"type": "selector",
						"value": "train"
					},
					{
						"dimension": "column_name",
						"type": "selector",
						"value": "addressState"
					}
				],
				"type": "and"
			},
			"granularity": "day",
			"intervals": {
				"intervals": [
					"2022-10-07T00:00:00Z/2022-10-14T00:00:00Z"
				],
				"type": "intervals"
			},
			"metric": {
				"metric": "count",
				"type": "numeric"
			},
			"queryType": "topN",
			"threshold": 100
		},
		"type": "query"
	},
	"granularity": "day",
	"intervals": {
		"intervals": [
			"2022-10-07T00:00:00Z/2022-10-14T00:00:00Z"
		],
		"type": "intervals"
	},
	"queryType": "timeseries"
}`
	location, _ := time.LoadLocation("UTC")
	start, _ := time.ParseInLocation(time.RFC3339,
		"2022-10-07T00:00:00Z",
		location)
	i := intervals.NewInterval()
	i.SetInterval(start, start.Add(time.Hour*24*7))
	interval := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})
	topN := query.NewTopN().
		SetThreshold(100).
		SetMetric(topnmetric.NewNumeric().SetMetric("count")).
		SetIntervals(interval).
		SetGranularity(granularity.NewSimple().SetGranularity(granularity.Day)).
		SetDimension(dimension.NewDefault().SetDimension("string_value")).
		SetDataSource(datasource2.NewTable().SetName("dc_94b4f5fdfde940979b79c50539d8322a_b42fde98efed4e638a0016b34b3c10cf_dataset_pre")).
		SetAggregations([]builder.Aggregator{aggregation.NewLongSum().SetName("count").SetFieldName("count")}).
		SetFilter(filter.NewAnd().SetFields([]builder.Filter{filter.NewSelector().SetDimension("_split_name_").SetValue("train"),
			filter.NewSelector().SetDimension("column_name").SetValue("addressState")}))
	subQuery := datasource2.NewQuery().SetQuery(topN)
	queryTest := query.NewTimeseries().SetGranularity(granularity.NewSimple().SetGranularity(granularity.Day)).
		SetIntervals(interval).SetDataSource(subQuery)

	queryTestJSON, err := json.Marshal(queryTest)
	assert.Nil(t, err)

	assert.JSONEq(t, expected, string(queryTestJSON))

}
