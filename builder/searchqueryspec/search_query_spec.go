package searchqueryspec

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

func Load(data []byte) (builder.SearchQuerySpec, error) {
	var s builder.SearchQuerySpec
	if string(data) == "null" {
		return s, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "all":
		s = NewAll()
	case "contains":
		s = NewContains()
	case "fragment":
		s = NewFragment()
	case "insensitiveContains":
		s = NewInsensitiveContains()
	case "regex":
		s = NewRegex()
	default:
		return nil, errors.New("unsupported searchqueryspec type")
	}
	return s, json.Unmarshal(data, &s)
}
