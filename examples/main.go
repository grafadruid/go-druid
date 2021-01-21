package main

import (
	"fmt"
	"github.com/grafadruid/go-druid/builder/intervals"
	"log"
	"time"

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

	i := intervals.NewInterval(time.Unix(0, 0), time.Now())
	is := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})

	c := aggregation.NewCount().SetName("count")
	aa := []builder.Aggregator{c}
	s := filter.NewSelector().SetDimension("countryName").SetValue("France")
	m := granularity.NewSimple().SetGranularity(granularity.All)
	ts := query.NewTimeseries().SetDataSource(t).SetIntervals(is).SetAggregations(aa).SetGranularity(m).SetFilter(s).SetLimit(10)
	var results interface{}
	d.Query().Execute(ts, &results)

	spew.Dump(results)
}
