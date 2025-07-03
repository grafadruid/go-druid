package havingspec

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

func Load(data []byte) (builder.HavingSpec, error) {
	var h builder.HavingSpec
	if string(data) == "null" {
		return h, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "always":
		h = NewAlways()
	case "and":
		h = NewAnd()
	case "dimSelector":
		h = NewDimSelector()
	case "equalTo":
		h = NewEqualTo()
	case "greaterThan":
		h = NewGreaterThan()
	case "lessThan":
		h = NewLessThan()
	case "never":
		h = NewNever()
	case "not":
		h = NewNot()
	case "or":
		h = NewOr()
	default:
		return nil, errors.New("unsupported havingspec type")
	}
	return h, json.Unmarshal(data, &h)
}
