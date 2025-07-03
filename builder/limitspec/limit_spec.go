package limitspec

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

func Load(data []byte) (builder.LimitSpec, error) {
	var l builder.LimitSpec
	if string(data) == "null" {
		return l, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "default":
		l = NewDefault()
	default:
		return nil, errors.New("unsupported limitspec type")
	}
	return l, json.Unmarshal(data, &l)
}
