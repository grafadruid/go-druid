//go:build mage
// +build mage

package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/h2oai/go-druid"
	"github.com/h2oai/go-druid/builder/query"
	"log"
)

func main() {
	d, err := druid.NewClient("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	var results []map[string]interface{}
	var param []query.SQLParameter
	param = append(param, query.NewSQLParameter("VARCHAR", "Salo Toraut"))
	param = append(param, query.NewSQLParameter("VARCHAR", "NB"))
	param = append(param, query.NewSQLParameter("VARCHAR", "false")) // BOOLEAN type fails the convent in api.
	param = append(param, query.NewSQLParameter("INTEGER", 31))
	context := make(map[string]interface{})
	context["sqlTimeZone"] = "America/Los_Angeles"
	query := query.NewSQL().SetQuery(
		`SELECT * FROM "wikipedia" WHERE page=? AND flags=? AND isUnpatrolled=? AND delta=?`).
		SetParameters(param).
		SetContext(context)
	_, err = d.Query().Execute(query, &results)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(results)
}
