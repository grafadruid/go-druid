package dimension

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Base struct {
	Type       string           `json:"type"`
	Dimension  string           `json:"dimension"`
	OutputName string           `json:"outputName"`
	OutputType types.OutputType `json:"outputType,omitempty"`
}

func NewBase() *Base {
	b := &Base{}
	return b
}

func (b *Base) SetType(typ string) *Base {
	b.Type = typ
	return b
}

func (b *Base) SetDimension(dimension string) *Base {
	b.Dimension = dimension
	return b
}

func (b *Base) SetOutputName(outputName string) *Base {
	b.OutputName = outputName
	return b
}

func (b *Base) SetOutputType(outputType types.OutputType) *Base {
	b.OutputType = outputType
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
	case "default":
		d = NewDefault()
	case "extraction":
		d = NewExtraction()
	case "listFiltered":
		d = NewListFiltered()
	case "lookup":
		d = NewLookup()
	case "prefixFiltered":
		d = NewPrefixFiltered()
	case "regexFiltered":
		d = NewRegexFiltered()
	}
	return d, json.Unmarshal(data, &d)
}
