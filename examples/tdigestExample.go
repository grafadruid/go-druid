//go:build mage
// +build mage

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/h2oai/go-druid"
	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/aggregation"
	"github.com/h2oai/go-druid/builder/datasource"
	"github.com/h2oai/go-druid/builder/granularity"
	"github.com/h2oai/go-druid/builder/intervals"
	"github.com/h2oai/go-druid/builder/postaggregation"
	"github.com/h2oai/go-druid/builder/query"
)

func getConnection() *druid.Client {
	d, err := druid.NewClient("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	status, _, err := d.Common().Status()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("{\"version\": \"" + status.Version + "\"}")
	return d
}

/*
	Prerequisite: These examples will work only if you have tdigest sketch data in your datastore.
	To experiment, you can use the doubles_sketch_data.tsv file attached in this repo. It is a copy of  https://github.com/apache/druid/blob/master/extensions-contrib/tdigestsketch/src/test/resources/doubles_sketch_data.tsv
*/

// tdigestSketchQuantilesPostAggUsingBuilder example using Builder Pattern
func tdigestSketchQuantilesPostAggUsingBuilder() {
	d := getConnection()
	table := datasource.NewTable().SetName("rollup-data")
	i := intervals.NewInterval()
	m := granularity.NewSimple().SetGranularity(granularity.All)
	i.SetInterval(time.Unix(0, 0), time.Now())
	is := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})

	// TDigest  Aggregation
	// valueTDS is the field holding the tdigest data in Your datastore
	// merged_sketch will hold the aggregated tdigest value
	atds := aggregation.NewTDigestSketch().SetName("merged_sketch").SetFieldName("valuesTDS")
	a := []builder.Aggregator{atds}

	//TDigest Post Aggregation Quantiles
	qf := postaggregation.NewQuantilesFromTDigestSketchField().
		SetType("fieldAccess").
		SetFieldName("merged_sketch")
	qa := postaggregation.NewQuantilesFromTDigestSketch().
		SetField(qf).
		SetFractions([]float64{0.25, 0.5, 0.75, 0.9, 0.95, 0.99}). // add additional quantiles as needed
		SetName("quantiles")
	pa := []builder.PostAggregator{qa}

	ts := query.NewTimeseries().SetDataSource(table).SetIntervals(is).SetAggregations(a).SetPostAggregations(pa).SetGranularity(m).SetLimit(10)
	var results interface{}
	_, err := d.Query().Execute(ts, &results)
	if err != nil {
		log.Fatalf("Execute failed, %s", err)
	}

	spew.Dump(results)
}

// tdigestSketchQuantilePostAggUsingBuilder example using Builder Pattern
func tdigestSketchQuantilePostAggUsingBuilder() {
	d := getConnection()
	table := datasource.NewTable().SetName("rollup-data")
	i := intervals.NewInterval()
	m := granularity.NewSimple().SetGranularity(granularity.All)
	i.SetInterval(time.Unix(0, 0), time.Now())
	is := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})

	// TDigest  Aggregation
	// valueTDS is the field holding the tdigest data in Your datastore
	// merged_sketch will hold the aggregated tdigest value
	atds := aggregation.NewTDigestSketch().SetName("merged_sketch").SetFieldName("valuesTDS")
	a := []builder.Aggregator{atds}

	//TDigest Post Aggregation Quantile
	qf := postaggregation.NewQuantileFromTDigestSketchField().
		SetType("fieldAccess").
		SetFieldName("merged_sketch")
	qa := postaggregation.NewQuantileFromTDigestSketch().
		SetField(qf).
		SetFraction(0.9).
		SetName("quantile")
	pa := []builder.PostAggregator{qa}

	ts := query.NewTimeseries().SetDataSource(table).SetIntervals(is).SetAggregations(a).SetPostAggregations(pa).SetGranularity(m).SetLimit(10)
	var results interface{}
	_, err := d.Query().Execute(ts, &results)
	if err != nil {
		log.Fatalf("Execute failed, %s", err)
	}

	spew.Dump(results)
}

// tdigestSketchUsingRuneQuery example using Native Query as the starting point
func main() {
	query := `{
			"queryType": "groupBy",
			"dataSource": {
			"type": "table",
			"name": "rollup-data1"
			},
			"granularity": "ALL",
			"dimensions": [],
			"aggregations": [{
				"type": "tDigestSketch",
				"name": "merged_sketch",
				"fieldName": "valuesTDS"
			}],
			"postAggregations": [{
				"type": "quantilesFromTDigestSketch",
				"name": "quantiles",
				"fractions": [0, 0.5, 0.9,1],
				"field": {
					"type": "fieldAccess",
					"fieldName": "merged_sketch"
				}
			}],
			"intervals": {
			"type": "intervals",
			"intervals": [
			  "-146136543-09-08T08:23:32.096Z/146140482-04-24T15:36:27.903Z"
			]
		  }
		}`

	d := getConnection()
	var results interface{}
	q, err := d.Query().Load([]byte(query))
	if err != nil {
		log.Fatalf("converting string to query object failed, %s", err)
	}

	resp, err := d.Query().Execute(q, &results)
	if err != nil {
		log.Fatalf("Execute failed, %s", err)
	}
	fmt.Printf("response code:%d", resp.StatusCode)
	spew.Dump(results)
}
