package extractionfn

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

func Load(data []byte) (builder.ExtractionFn, error) {
	var e builder.ExtractionFn
	if string(data) == "null" {
		return e, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "bucket":
		e = NewBucket()
	case "cascade":
		e = NewCascade()
	case "identity":
		e = NewIdentity()
	case "javascript":
		e = NewJavascript()
	case "lookup":
		e = NewLookup()
	case "lower":
		e = NewLower()
	case "partial":
		e = NewPartial()
	case "regex":
		e = NewRegex()
	case "registeredLookup":
		e = NewRegisteredLookup()
	case "searchQuery":
		e = NewSearchQuery()
	case "stringFormat":
		e = NewStringFormat()
	case "strlen":
		e = NewStrlen()
	case "substring":
		e = NewSubstring()
	case "time":
		e = NewTime()
	case "timeFormat":
		e = NewTimeFormat()
	case "upper":
		e = NewUpper()
	default:
		return nil, errors.New("unsupported extractionfn type")
	}
	return e, json.Unmarshal(data, &e)
}
