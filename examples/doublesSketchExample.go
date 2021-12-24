//+build mage

package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/grafadruid/go-druid"
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/aggregation"
	"github.com/grafadruid/go-druid/builder/datasource"
	"github.com/grafadruid/go-druid/builder/dimension"
	"github.com/grafadruid/go-druid/builder/granularity"
	"github.com/grafadruid/go-druid/builder/intervals"
	"github.com/grafadruid/go-druid/builder/postaggregation"
	"github.com/grafadruid/go-druid/builder/query"
	"log"
	"time"
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
	Prerequisite: These examples will work only if you have doubles sketch data in your datastore.
	To experiment, you can use the doubles_sketch_data.tsv file attached in this repo. It is a copy of  https://github.com/apache/druid/blob/master/extensions-contrib/tdigestsketch/src/test/resources/doubles_sketch_data.tsv
*/

// doublesSketchUsingBuilder example using Builder Pattern
func doublesSketchUsingBuilder() {
	d := getConnection()
	table := datasource.NewTable().SetName("double-sketch")
	i := intervals.NewInterval()
	m := granularity.NewSimple().SetGranularity(granularity.All)
	i.SetInterval(time.Unix(0, 0), time.Now())
	is := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})

	ads := aggregation.NewQuantilesDoublesSketch().SetName("a1:agg").SetFieldName("latencySketch").SetK(128)
	a := []builder.Aggregator{ads}

	//TDigest Post Aggregation
	qf := postaggregation.NewQuantilesDoublesSketchToQuantileField().
		SetType("fieldAccess").
		SetFieldName("a1:agg").
		SetName("tp90")
	qa := postaggregation.NewQuantilesDoublesSketchToQuantile().
		SetField(qf).
		SetFraction(0.90). // add additional quantiles as needed
		SetName("tp90")
	pa := []builder.PostAggregator{qa}

	ts := query.NewTimeseries().SetDataSource(table).SetIntervals(is).SetAggregations(a).SetPostAggregations(pa).SetGranularity(m).SetLimit(10)
	di := dimension.NewDefault().SetDimension("svcAssetId")
	da := []builder.Dimension{di}
	gb := query.NewGroupBy().SetDimensions(da).SetDataSource(table).SetIntervals(is).
		SetAggregations(a).
		SetPostAggregations(pa).
		SetGranularity(m)

	tsJSON, _ := json.Marshal(ts)
	log.Printf(string(tsJSON))

	gbJSON, _ := json.Marshal(gb)
	log.Printf(string(gbJSON))

	var results interface{}
	_, err := d.Query().Execute(ts, &results)
	if err != nil {
		log.Fatalf("Execute failed, %s", err)
	}

	spew.Dump(results)
}

// doublesSketchUsingRuneQuery example using Native Query as the starting point
func main() {
	query := `{
			"queryType": "groupBy",
			"dataSource": {
				"type": "table",
				"name": "double-sketch"
			},
			"granularity": "ALL",
			"dimensions": [{
			  "type": "default",
			  "dimension": "uniqueId"
    		}],
			"aggregations": [{
				"type": "quantilesDoublesSketch",
				"name": "a1:agg",
				"fieldName": "latencySketch",
				"k": 128
			}],
			"postAggregations": [{
				"type": "quantilesDoublesSketchToQuantile",
				"name": "tp90",
				"fraction": 0.90,
				"field": {
					"type": "fieldAccess",
					"name": "tp90",
					"fieldName": "a1:agg"
				}
			}],
			"intervals": {
			"type": "intervals",
			"intervals": [
			  "-146136543-09-08T08:23:32.096Z/146140482-04-24T15:36:27.903Z"
			]
		  }
		}`

	doublesSketchUsingBuilder()

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
