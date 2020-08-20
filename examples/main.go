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
}
