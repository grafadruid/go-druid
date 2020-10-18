package havingspec

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Type string `json:"type"`
}

func NewBase() *Base {
	b := &Base{}
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Type = typ
	return b
}

func Load(data []byte) (query.Dimension, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d query.Dimension
	switch t.Typ {
	case "always":
		d = NewAlways()
	case "and":
		d = NewAnd()
	case "dimSelector":
		d = NewDimSelector()
	case "equalTo":
		d = NewEqualTo()
	case "greaterThan":
		d = NewGreaterThan()
	case "lessThan":
		d = NewLessThan()
	case "never":
		d = NewNever()
	case "not":
		d = NewNot()
	case "or":
		d = NewOr()
	}
	return d, json.Unmarshal(data, &d)
}
