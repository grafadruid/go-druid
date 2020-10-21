package topnmetric

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Typ string `json:"type"`
}

func NewBase() *Base {
	b := &Base{}
	b.SetType("base")
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() query.ComponentType {
	return b.Typ
}

func Load(data []byte) (query.TopNMetric, error) {
	var t struct {
		Typ string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var tnm query.TopNMetric
	switch t.Typ {
	case "alphaNumeric":
		tnm = NewAlphaNumeric()
	case "dimension":
		tnm = NewDimension()
	case "inverted":
		tnm = NewInverted()
	case "lexicographic":
		tnm = NewLexicographic()
	case "numeric":
		tnm = NewNumeric()
	}
	return tnm, json.Unmarshal(data, &tnm)
}
