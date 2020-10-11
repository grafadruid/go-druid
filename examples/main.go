package main

import (
	"fmt"
	"log"

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
	fmt.Println(status.Version)
	q, err := d.Query().Load([]byte("{\"queryType\":\"scan\",\"dataSource\":{\"type\":\"table\",\"name\":\"wikipedia\"},\"intervals\":[\"1980-06-12T22:30:00.000Z/2020-01-26T23:00:00.000Z\"],\"virtualColumns\":[{\"type\":\"expression\",\"name\":\"v0\",\"expression\":\"'France'\",\"outputType\":\"STRING\"}],\"resultFormat\":\"compactedList\",\"batchSize\":20480,\"limit\":10,\"order\":\"none\",\"filter\":{\"type\":\"selector\",\"dimension\":\"countryName\",\"value\":\"France\",\"extractionFn\":null},\"columns\":[\"__time\",\"channel\",\"cityName\",\"comment\",\"count\",\"countryIsoCode\",\"diffUrl\",\"flags\",\"isAnonymous\",\"isMinor\",\"isNew\",\"isRobot\",\"isUnpatrolled\",\"metroCode\",\"namespace\",\"page\",\"regionIsoCode\",\"regionName\",\"sum_added\",\"sum_commentLength\",\"sum_deleted\",\"sum_delta\",\"sum_deltaBucket\",\"user\",\"v0\"],\"legacy\":false,\"context\":{\"sqlOuterLimit\":100,\"sqlQueryId\":\"b12ac7bb-7cc5-4873-b19d-1cd95264e01b\"},\"descending\":false,\"granularity\":{\"type\":\"all\"}}"))
	var results []map[string]interface{}
	d.Query().Execute(q, &results)
	for _, evt := range results[0]["events"].([]interface{}) {
		fmt.Println("********************************ROW********************************")
		fmt.Println(evt)
	}
}
