package havingspec

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.Dimension, error) {
	var t struct {
		Typ builder.ComponentType `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var d builder.Dimension
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
