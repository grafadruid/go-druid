package extractionfn

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Typ string `json:"type"`
}

func NewBase() *Base {
	b := &Base{}
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() query.ComponentType {
	return b.Typ
}

func Load(data []byte) (query.ExtractionFn, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var e query.ExtractionFn
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
	}
	return e, json.Unmarshal(data, &e)
}
