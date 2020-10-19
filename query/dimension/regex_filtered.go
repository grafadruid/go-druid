package dimension

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type RegexFiltered struct {
	Base
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

func (r *RegexFiltered) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Delegate json.RawMessage `json:"delegate"`
		Pattern  string          `json:"pattern"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	d, err := Load(tmp.Delegate)
	if err != nil {
		return err
	}
	r.Base = tmp.Base
	r.Delegate = d
	r.Pattern = tmp.Pattern
	return nil
}
