package bound

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

func Load(data []byte) (builder.Bound, error) {
	var b builder.Bound
	if string(data) == "null" {
		return b, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "polygon":
		b = NewPolygon()
	case "radius":
		b = NewRadius()
	case "rectangular":
		b = NewRectangular()
	default:
		return nil, errors.New("unsupported bound type")
	}
	return b, json.Unmarshal(data, &b)
}
