package dimension

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/types"
)

type RegexFiltered struct {
	Base
	Delegate builder.Dimension `json:"delegate,omitempty"`
	Pattern  string            `json:"pattern,omitempty"`
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

func (r *RegexFiltered) SetDelegate(delegate builder.Dimension) *RegexFiltered {
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
		Delegate json.RawMessage `json:"delegate,omitempty"`
		Pattern  string          `json:"pattern,omitempty"`
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
