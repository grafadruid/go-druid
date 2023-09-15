package dimension

import "github.com/h2oai/go-druid/builder/types"

type Default struct {
	Base
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
