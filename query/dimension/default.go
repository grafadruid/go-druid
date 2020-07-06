package dimension

import "github.com/grafadruid/go-druid/query/types"

type Def struct {
	*Base
}

func NewDef() *Def {
	d := &Def{}
	d.SetType("def")
	return d
}

func (d *Def) SetDimension(dimension string) *Def {
	d.Base.SetDimension(dimension)
	return d
}

func (d *Def) SetOutputName(outputName string) *Def {
	d.Base.SetOutputName(outputName)
	return d
}

func (d *Def) SetOutputType(outputType types.OutputType) *Def {
	d.Base.SetOutputType(outputType)
	return d
}
