package topnmetric

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Base struct {
	Type string `json:"type"`
}

func NewBase() *Base {
	b := &Base{}
	b.SetType("base")
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Type = typ
	return b
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
