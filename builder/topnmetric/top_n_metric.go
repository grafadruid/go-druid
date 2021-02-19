package topnmetric

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type,omitempty"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.TopNMetric, error) {
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var tnm builder.TopNMetric
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
	default:
		return nil, errors.New("unsupported type")
	}
	return tnm, json.Unmarshal(data, &tnm)
}
