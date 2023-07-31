package main

import (
	"log"
	"time"

	"github.com/adjoeio/go-druid"
	"github.com/adjoeio/go-druid/builder"
	"github.com/adjoeio/go-druid/builder/aggregation"
	datasource2 "github.com/adjoeio/go-druid/builder/datasource"
	"github.com/adjoeio/go-druid/builder/dimension"
	"github.com/adjoeio/go-druid/builder/extractionfn"
	"github.com/adjoeio/go-druid/builder/filter"
	"github.com/adjoeio/go-druid/builder/granularity"
	"github.com/adjoeio/go-druid/builder/intervals"
	"github.com/adjoeio/go-druid/builder/query"
	"github.com/adjoeio/go-druid/builder/topnmetric"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	d, err := druid.NewClient("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{}

	_, err = d.Query().Execute(MakeScanSubQuery(), &results)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(results)

	_, err = d.Query().Execute(MakeTimeSeriesSubQuery(), &results)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(results)
}

// MakeScanSubQuery Result is nil
func MakeScanSubQuery() *query.Scan {
	location, _ := time.LoadLocation("UTC")

	start, _ := time.ParseInLocation(time.RFC3339Nano,
		"1980-06-12T22:30:00.000Z",
		location)
	end, _ := time.ParseInLocation(time.RFC3339Nano,
		"2023-01-26T23:00:00.000Z",
		location)
	i := intervals.NewInterval()
	i.SetInterval(start, end)
	interval := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})

	selectorExtractionFn := extractionfn.NewLower().SetLocale("")
	a := filter.NewSelector().SetDimension("countryName").SetValue("france").
		SetExtractionFn(selectorExtractionFn)
	query.NewSegmentMetadata().SetIntervals(interval)
	datasourceSubQuery := query.NewScan().
		SetDataSource(datasource2.NewTable().SetName("A")).SetColumns([]string{"AT"}).
		SetIntervals(interval)
	datasource := datasource2.NewQuery().SetQuery(datasourceSubQuery)
	scan := query.NewScan()
	queryTest := scan.SetLimit(10).SetIntervals(interval).SetFilter(a).
		SetColumns([]string{
			"__time", "channel", "cityName", "comment", "count", "countryIsoCode",
			"diffUrl", "flags", "isAnonymous", "isMinor", "isNew", "isRobot", "isUnpatrolled",
			"metroCode", "namespace", "page", "regionIsoCode", "regionName", "sum_added",
			"sum_commentLength", "sum_deleted", "sum_delta", "sum_deltaBucket", "user",
		}).
		SetBatchSize(20480).
		SetDataSource(datasource).SetOrder(query.None)
	return queryTest
}

// MakeTimeSeriesSubQuery Results are timestamp 2022-10-07 ~ 13
func MakeTimeSeriesSubQuery() *query.Timeseries {
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
		SetFilter(filter.NewAnd().SetFields([]builder.Filter{
			filter.NewSelector().SetDimension("_split_name_").SetValue("train"),
			filter.NewSelector().SetDimension("column_name").SetValue("addressState"),
		}))
	subQuery := datasource2.NewQuery().SetQuery(topN)
	queryTest := query.NewTimeseries().SetGranularity(granularity.NewSimple().SetGranularity(granularity.Day)).
		SetIntervals(interval).SetDataSource(subQuery)
	return queryTest
}
