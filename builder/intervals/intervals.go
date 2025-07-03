package intervals

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

func Load(data []byte) (builder.Intervals, error) {
	var i builder.Intervals
	if string(data) == "null" {
		return i, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "intervals":
		i = NewIntervals()
	default:
		return nil, errors.New("unsupported intervals type")
	}
	return i, json.Unmarshal(data, &i)
}
