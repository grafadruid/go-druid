package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/grafadruid/go-druid"
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

	var results interface{}
	q, err := d.Query().Load([]byte("{\"queryType\":\"scan\",\"dataSource\":{\"type\":\"table\",\"name\":\"wikipedia\"},\"intervals\":[\"1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z\"],\"virtualColumns\":[{\"type\":\"expression\",\"name\":\"v0\",\"expression\":\"'France'\",\"outputType\":\"STRING\"}],\"resultFormat\":\"compactedList\",\"batchSize\":20480,\"limit\":1,\"order\":\"none\",\"filter\":{\"type\":\"selector\",\"dimension\":\"countryName\",\"value\":\"France\",\"extractionFn\":null},\"columns\":[\"__time\",\"channel\",\"cityName\",\"comment\",\"count\",\"countryIsoCode\",\"diffUrl\",\"flags\",\"isAnonymous\",\"isMinor\",\"isNew\",\"isRobot\",\"isUnpatrolled\",\"metroCode\",\"namespace\",\"page\",\"regionIsoCode\",\"regionName\",\"sum_added\",\"sum_commentLength\",\"sum_deleted\",\"sum_delta\",\"sum_deltaBucket\",\"user\",\"v0\"],\"legacy\":false,\"context\":{\"sqlOuterLimit\":100,\"sqlQueryId\":\"b12ac7bb-7cc5-4873-b19d-1cd95264e01b\"},\"descending\":false,\"granularity\":{\"type\":\"all\"}}"))
	spew.Dump(q)
	d.Query().Execute(q, &results)
	spew.Dump(results)

	q, err = d.Query().Load([]byte("{\"batchSize\":20480,\"columns\":[\"__time\",\"channel\",\"cityName\",\"comment\",\"count\",\"countryIsoCode\",\"diffUrl\",\"flags\",\"isAnonymous\",\"isMinor\",\"isNew\",\"isRobot\",\"isUnpatrolled\",\"metroCode\",\"namespace\",\"page\",\"regionIsoCode\",\"regionName\",\"sum_added\",\"sum_commentLength\",\"sum_deleted\",\"sum_delta\",\"sum_deltaBucket\",\"user\"],\"dataSource\":{\"name\":\"wikipedia\",\"type\":\"table\"},\"filter\":{\"dimension\":\"countryName\",\"extractionFn\":{\"locale\":\"\",\"type\":\"lower\"},\"type\":\"selector\",\"value\":\"france\"},\"intervals\":[\"1980-01-26T23:00:00.000Z/2020-01-26T23:00:00.000Z\"],\"limit\":10,\"order\":\"descending\",\"queryType\":\"scan\", \"resultFormat\":\"compactedList\"}"))
	spew.Dump(q)
	d.Query().Execute(q, &results)
	spew.Dump(results)

	q, err = d.Query().Load([]byte("{\"batchSize\":20480,\"columns\":[\"__time\",\"channel\",\"cityName\",\"comment\",\"count\",\"countryIsoCode\",\"diffUrl\",\"flags\",\"isAnonymous\",\"isMinor\",\"isNew\",\"isRobot\",\"isUnpatrolled\",\"metroCode\",\"namespace\",\"page\",\"regionIsoCode\",\"regionName\",\"sum_added\",\"sum_commentLength\",\"sum_deleted\",\"sum_delta\",\"sum_deltaBucket\",\"user\"],\"dataSource\":{\"type\":\"query\",\"query\":{\"queryType\":\"scan\",\"dataSource\":{\"type\":\"table\",\"name\":\"A\"},\"columns\":[\"AT\"],\"intervals\":[\"1980-01-26T23:00:00.000Z/2020-01-26T23:00:00.000Z\"]}},\"filter\":{\"dimension\":\"countryName\",\"extractionFn\":{\"locale\":\"\",\"type\":\"lower\"},\"type\":\"selector\",\"value\":\"france\"},\"intervals\":[\"1980-01-26T23:00:00.000Z/2020-01-26T23:00:00.000Z\"],\"limit\":10,\"order\":\"descending\",\"queryType\":\"scan\"}"))
	spew.Dump(q, err)
	d.Query().Execute(q, &results)
	spew.Dump(results)

	q, err = d.Query().Load([]byte("{\"context\":{\"con\":\"text\"},\"query\":\"SELECT \\\"__time\\\", \\\"channel\\\", \\\"cityName\\\", \\\"comment\\\", \\\"count\\\", \\\"countryIsoCode\\\", \\\"countryName\\\", \\\"diffUrl\\\", \\\"flags\\\", \\\"isAnonymous\\\", \\\"isMinor\\\", \\\"isNew\\\", \\\"isRobot\\\", \\\"isUnpatrolled\\\", \\\"metroCode\\\", \\\"namespace\\\", \\\"page\\\", \\\"regionIsoCode\\\", \\\"regionName\\\", \\\"sum_added\\\", \\\"sum_commentLength\\\", \\\"sum_deleted\\\", \\\"sum_delta\\\", \\\"sum_deltaBucket\\\", \\\"user\\\"\\nFROM \\\"wikipedia\\\"\\nWHERE \\\"countryName\\\" = 'France'\\nLIMIT 10\",\"queryType\":\"sql\", \"resultFormat\":\"array\", \"header\": true}"))
	spew.Dump(q, err)
	d.Query().Execute(q, &results)
	spew.Dump(results)
}
