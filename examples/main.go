package main

import (
	"fmt"
	"log"
	"time"

	"github.com/grafadruid/go-druid/builder/intervals"

	"github.com/davecgh/go-spew/spew"
	"github.com/grafadruid/go-druid"
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/aggregation"
	"github.com/grafadruid/go-druid/builder/datasource"
	"github.com/grafadruid/go-druid/builder/filter"
	"github.com/grafadruid/go-druid/builder/granularity"
	"github.com/grafadruid/go-druid/builder/query"
)

func main() {
	d, err := druid.NewClient("http://localhost:8082")
	if err != nil {
		log.Fatal(err)
	}
	status, _, err := d.Common().Status()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("{\"version\": \"" + status.Version + "\"}")

	t := datasource.NewTable().SetName("wikipedia")

	i := intervals.NewInterval()
	i.SetInterval(time.Unix(0, 0), time.Now())
	i2 := intervals.NewInterval()
	i2.SetIntervalWithString("2021-01-21T14:59:05.000Z", "P1D")
	is := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i, i2})

	c := aggregation.NewCount().SetName("count")
	aa := []builder.Aggregator{c}
	s := filter.NewSelector().SetDimension("countryName").SetValue("France")
	m := granularity.NewSimple().SetGranularity(granularity.All)
	ts := query.NewTimeseries().SetDataSource(t).SetIntervals(is).SetAggregations(aa).SetGranularity(m).SetFilter(s).SetLimit(10)
	var results interface{}
	_, err = d.Query().Execute(ts, &results)
	if err != nil {
		log.Fatalf("Execute failed, %s", err)
	}

	spew.Dump(results)
}
