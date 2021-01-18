package main

import (
	"fmt"
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
	"github.com/grafadruid/go-druid/builder/types"
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
	i := types.NewInterval(time.Unix(0, 0), time.Now())
	c := aggregation.NewCount().SetName("count")
	aa := []builder.Aggregator{c}
	s := filter.NewSelector().SetDimension("country").SetValue("France")
	m := granularity.NewSimple().SetGranularity(granularity.Minute)
	ts := query.NewTimeseries().SetDataSource(t).SetIntervals([]types.Interval{i}).SetAggregations(aa).SetGranularity(m).SetFilter(s).SetLimit(10)
	var results interface{}
	d.Query().Execute(ts, &results)
	spew.Dump(results)
}
