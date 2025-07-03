package datasource

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/common"
)

// Base embeds the shared TypeOnlyBase to eliminate code duplication
type Base struct {
	common.TypeOnlyBase
}

func Load(data []byte) (builder.DataSource, error) {
	var d builder.DataSource
	if string(data) == "null" {
		return d, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "globalTable":
		d = NewGlobalTable()
	case "inline":
		d = NewInline()
	case "join":
		d = NewJoin()
	case "lookup":
		d = NewLookup()
	case "query":
		d = NewQuery()
	case "table":
		d = NewTable()
	case "union":
		d = NewUnion()
	default:
		return nil, errors.New("unsupported datasource type")
	}
	return d, json.Unmarshal(data, &d)
}
