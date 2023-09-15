//go:build mage
// +build mage

package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/h2oai/go-druid"
	"github.com/h2oai/go-druid/builder"
)

func loadAndExecute(d *druid.Client, qry []byte) {
	var err error
	var q builder.Query
	var results interface{}
	q, err = d.Query().Load(qry)
	if err != nil {
		log.Fatalf("Load failed, %s (query: %s)", err, string(qry))
	}
	spew.Dump(q)
	_, err = d.Query().Execute(q, &results)
	if err != nil {
		log.Fatalf("Execute failed, %s (query: %s)", err, string(qry))
	}
	spew.Dump(results)
}

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

	var qry string

	qry = `{"queryType":"groupBy","dataSource":{"type":"table","name":"wikipedia"},"intervals":{"type":"intervals","intervals":["-146136543-09-08T08:23:32.096Z/146140482-04-24T15:36:27.903Z"]},"virtualColumns":[{"type":"expression","name":"v0","expression":"'France'","outputType":"STRING"}],"filter":{"type":"selector","dimension":"countryName","value":"France","extractionFn":null},"granularity":{"type":"all"},"dimensions":[{"type":"default","dimension":"v0","outputName":"d0","outputType":"STRING"},{"type":"default","dimension":"__time","outputName":"d1","outputType":"LONG"}],"aggregations":[{"type":"longSum","name":"a0","fieldName":"count","expression":null}],"postAggregations":[{"type":"expression","name":"s0","expression":"'France'","ordering":null}],"having":null,"limitSpec":{"type":"default","columns":[],"limit":100},"context":{"sqlOuterLimit":100,"sqlQueryId":"95c6f52f-9dcf-421c-94c7-c50fabbc16b0","useApproximateCountDistinct":false,"useApproximateTopN":false},"descending":false}`
	loadAndExecute(d, []byte(qry))

	qry = `{"queryType":"scan","dataSource":{"type":"table","name":"wikipedia"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"virtualColumns":[{"type":"expression","name":"v0","expression":"'France'","outputType":"STRING"}],"resultFormat":"compactedList","batchSize":20480,"limit":1,"order":"none","filter":{"type":"selector","dimension":"countryName","value":"France","extractionFn":null},"columns":["__time","channel","cityName","comment","count","countryIsoCode","diffUrl","flags","isAnonymous","isMinor","isNew","isRobot","isUnpatrolled","metroCode","namespace","page","regionIsoCode","regionName","sum_added","sum_commentLength","sum_deleted","sum_delta","sum_deltaBucket","user","v0"],"legacy":false,"context":{"sqlOuterLimit":100,"sqlQueryId":"b12ac7bb-7cc5-4873-b19d-1cd95264e01b"},"descending":false}`
	loadAndExecute(d, []byte(qry))

	qry = `{"batchSize":20480,"columns":["__time","channel","cityName","comment","count","countryIsoCode","diffUrl","flags","isAnonymous","isMinor","isNew","isRobot","isUnpatrolled","metroCode","namespace","page","regionIsoCode","regionName","sum_added","sum_commentLength","sum_deleted","sum_delta","sum_deltaBucket","user","v0"],"context":{"plopa":"plep"},"dataSource":{"name":"wikipedia","type":"table"},"filter":{"dimension":"countryName","type":"selector","value":"France"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"limit":10,"order":"none","queryType":"scan","virtualColumns":[{"expression":"'France'","name":"v0","outputType":"STRING","type":"expression"}]}`
	loadAndExecute(d, []byte(qry))

	qry = `{"batchSize":20480,"columns":["__time","channel","cityName","comment","count","countryIsoCode","diffUrl","flags","isAnonymous","isMinor","isNew","isRobot","isUnpatrolled","metroCode","namespace","page","regionIsoCode","regionName","sum_added","sum_commentLength","sum_deleted","sum_delta","sum_deltaBucket","user"],"dataSource":{"name":"wikipedia","type":"table"},"filter":{"dimension":"countryName","extractionFn":{"locale":"","type":"lower"},"type":"selector","value":"france"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"limit":10,"order":"descending","queryType":"scan", "resultFormat":"compactedList"}`
	loadAndExecute(d, []byte(qry))

	// https://github.com/grafadruid/go-druid/issues/15
	qry = `{"batchSize":20480,"columns":["__time","channel","cityName","comment","count","countryIsoCode","diffUrl","flags","isAnonymous","isMinor","isNew","isRobot","isUnpatrolled","metroCode","namespace","page","regionIsoCode","regionName","sum_added","sum_commentLength","sum_deleted","sum_delta","sum_deltaBucket","user"],"dataSource":{"type":"query","query":{"queryType":"scan","dataSource":{"type":"table","name":"A"},"columns":["AT"],"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]}}},"filter":{"dimension":"countryName","extractionFn":{"locale":"","type":"lower"},"type":"selector","value":"france"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"limit":10,"order":"descending","queryType":"scan"}`
	//loadAndExecute(d, []byte(qry))

	qry = `{"context":{"con":"text"},"query":"SELECT * \nFROM \"wikipedia\"\nWHERE \"countryName\" = 'France'\nLIMIT 10","queryType":"sql", "resultFormat":"array", "header": true}`
	loadAndExecute(d, []byte(qry))

	qry = `{"aggregations":[{"name":"count","type":"count"}],"context":{"plop":"plep"},"dataSource":{"name":"wikipedia","type":"table"},"filter":{"dimension":"countryName","extractionFn":null,"type":"selector","value":"France"},"granularity":"minute","intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"limit":10,"postAggregations":[],"queryType":"timeseries","virtualColumns":[]}`
	loadAndExecute(d, []byte(qry))

	qry = `{"aggregations":[{"name":"count","type":"count"}],"context":{"plop":"plep"},"dataSource":{"name":"wikipedia","type":"table"},"filter":{"dimension":"countryName","extractionFn":null,"type":"selector","value":"France"},"granularity":{"duration":1000,"origin":"2016-06-27T13:30:00.000Z","type":"duration"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"limit":10,"postAggregations":[],"queryType":"timeseries","virtualColumns":[]}`
	loadAndExecute(d, []byte(qry))

	qry = `{"aggregations":[{"fieldName":"sum_delta","name":"delta","type":"longSum"}],"context":{"plop":"plep"},"dataSource":{"name":"wikipedia","type":"table"},"dimension":{"dimension":"regionName","outputName":"region","outputType":"STRING","type":"default"},"filter":{"dimension":"countryName","extractionFn":null,"type":"selector","value":"France"},"granularity":"hour","intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"metric":{"metric":"delta","type":"numeric"},"postAggregations":[],"queryType":"topN","threshold":50,"virtualColumns":[]}`
	loadAndExecute(d, []byte(qry))

	qry = `{"aggregations":[{"fieldName":"sum_delta","name":"delta","type":"longSum"}],"context":{"plop":"plep"},"dataSource":{"name":"wikipedia","type":"table"},"dimensions":[{"dimension":"countryName","outputName":"country","outputType":"STRING","type":"default"},{"dimension":"regionName","outputName":"region","outputType":"STRING","type":"default"}],"filter":{"dimension":"countryName","extractionFn":null,"type":"selector","value":"France"},"granularity":"minute","intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"postAggregations":[],"queryType":"groupBy","subtotalsSpec":[],"virtualColumns":[]}`
	loadAndExecute(d, []byte(qry))

	qry = `{"context":{"plop":"plep"},"dataSource":{"name":"wikipedia","type":"table"},"filter":{"dimension":"countryName","extractionFn":null,"type":"selector","value":"France"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"limit":0,"query":{"type":"contains","value":"Allier"},"queryType":"search","searchDimensions":[{"dimension":"regionName","outputType":"STRING","type":"default"}]}`
	loadAndExecute(d, []byte(qry))

	qry = `{"bound":"minTime","context":{},"dataSource":{"name":"wikipedia","type":"table"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"queryType":"timeBoundary"}`
	loadAndExecute(d, []byte(qry))

	qry = `{"context":{"a":"a"},"dataSource":{"name":"wikipedia","type":"table"},"queryType":"dataSourceMetadata"}`
	loadAndExecute(d, []byte(qry))

	qry = `{"analysisTypes":["minmax","interval","size","cardinality","timestampSpec","queryGranularity","aggregators","rollup"],"context":{"":""},"dataSource":{"name":"wikipedia","type":"table"},"intervals":{"type":"intervals","intervals":["1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z"]},"queryType":"segmentMetadata"}`
	loadAndExecute(d, []byte(qry))
}
