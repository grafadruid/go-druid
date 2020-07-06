package dimension

import "github.com/grafadruid/go-druid/query/types"

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
