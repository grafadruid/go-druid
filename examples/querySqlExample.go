//go:build mage
// +build mage

package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/grafadruid/go-druid"
	"github.com/grafadruid/go-druid/builder/query"
	"log"
)

func main() {
	d, err := druid.NewClient("http://localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	var results []map[string]interface{}
	var param []query.SQLParameter
	pageParam := query.NewSQLParameter("VARCHAR", "Salo Toraut")
	flagParam := query.NewSQLParameter("VARCHAR", "NB")
	isUnpatrolledParam := query.NewSQLParameter("VARCHAR", "false") // BOOLEAN type fails the convent in api.
	deltaParam := query.NewSQLParameter("INTEGER", 31)              // This is why I changed the type of values to interface.
	param = append(param, pageParam)
	param = append(param, flagParam)
	param = append(param, isUnpatrolledParam)
	param = append(param, deltaParam)
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
