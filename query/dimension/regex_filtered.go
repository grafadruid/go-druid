package dimension

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type RegexFiltered struct {
	*Base
	Delegate query.Dimension `json:"delegate"`
	Pattern  string          `json:"pattern"`
}

func NewRegexFiltered() *RegexFiltered {
	r := &RegexFiltered{}
	r.SetType("regexFiltered")
	return r
}

func (r *RegexFiltered) SetDimension(dimension string) *RegexFiltered {
	r.Base.SetDimension(dimension)
	return r
}

func (r *RegexFiltered) SetOutputName(outputName string) *RegexFiltered {
	r.Base.SetOutputName(outputName)
	return r
}

func (r *RegexFiltered) SetOutputType(outputType types.OutputType) *RegexFiltered {
	r.Base.SetOutputType(outputType)
	return r
}

func (r *RegexFiltered) SetDelegate(delegate query.Dimension) *RegexFiltered {
	r.Delegate = delegate
	return r
}

func (r *RegexFiltered) SetPattern(pattern string) *RegexFiltered {
	r.Pattern = pattern
	return r
}
