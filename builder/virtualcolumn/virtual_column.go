package virtualcolumn

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

func Load(data []byte) (builder.VirtualColumn, error) {
	var v builder.VirtualColumn
	if string(data) == "null" {
		return v, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "expression":
		v = NewExpression()
	default:
		return nil, errors.New("unsupported virtualcolumn type")
	}
	return v, json.Unmarshal(data, &v)
}
