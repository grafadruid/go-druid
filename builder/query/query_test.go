package query

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectSQL(t *testing.T) {
	var param []SQLParameter
	param = append(param, NewSQLParameter("VARCHAR", "Salo Toraut"))
	param = append(param, NewSQLParameter("VARCHAR", "NB"))
	param = append(param, NewSQLParameter("BOOLEAN", "false"))
	param = append(param, NewSQLParameter("INTEGER", 31)) // This is why I changed the type of values to interface.
	param = append(param, NewSQLParameter("TIMESTAMP", "2016-06-27T00:00:11.080Z"))
	query := NewSQL().SetQuery(
		`SELECT * FROM "wikipedia" WHERE page=? AND flags=? AND isUnpatrolled=?  AND delta=? AND __time=?`).SetParameters(param)
	expected := `{
  "queryType": "sql",
  "query": "SELECT * FROM \"wikipedia\" WHERE page=? AND flags=? AND isUnpatrolled=?  AND delta=? AND __time=?",
  "parameters": [
    {
      "type": "VARCHAR",
      "value": "Salo Toraut"
    },
    {
      "type": "VARCHAR",
      "value": "NB"
    },
    {
      "type": "BOOLEAN",
      "value": "false"
    },
    {
      "type": "INTEGER",
      "value": 31
    },
    {
      "type": "TIMESTAMP",
      "value": "2016-06-27T00:00:11.080Z"
    }
  ]
}`
	expressionJSON, err := json.Marshal(query)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(expressionJSON))
}

func TestAllSetterSQL(t *testing.T) {
	var param []SQLParameter
	param = append(param, NewSQLParameter("VARCHAR", "Salo Toraut"))
	context := make(map[string]interface{})
	context["sqlTimeZone"] = "America/Los_Angeles"

	query := NewSQL().SetQuery(`select * from "wikipedia" WHERE page=?`).
		SetParameters(param).
		SetHeader(true).
		SetTypesHeader(true).
		SetResultFormat("array").
		SetSQLTypesHeader(true).
		SetContext(context)
	expected := `{
  "queryType": "sql",
  "context": {
    "sqlTimeZone": "America/Los_Angeles"
  },
  "query": "select * from \"wikipedia\" WHERE page=?",
  "resultFormat": "array",
  "header": true,
  "typesHeader": true,
  "sqlTypesHeader": true,
  "parameters": [
    {
      "type": "VARCHAR",
      "value": "Salo Toraut"
    }
  ]
}`
	expressionJSON, err := json.Marshal(query)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(expressionJSON))

}
