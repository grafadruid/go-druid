package dimension

import (
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/types"
)

type Default struct {
	Base
}

func (d *Default) GetDimension() string {
	return d.Base.Dimension
}

func (d *Default) GetOutputName() string {
	return d.Base.OutputName
}

func (d *Default) GetOutputType() types.OutputType {
	return d.Base.OutputType
}

func (d *Default) GetExtractionFn() builder.ExtractionFn {
	return nil
}

func NewDefault() *Default {
	d := &Default{}
	d.SetType("default")
	return d
}

func (d *Default) SetDimension(dimension string) *Default {
	d.Base.SetDimension(dimension)
	return d
}

func (d *Default) SetOutputName(outputName string) *Default {
	d.Base.SetOutputName(outputName)
	return d
}

func (d *Default) SetOutputType(outputType types.OutputType) *Default {
	d.Base.SetOutputType(outputType)
	return d
}
