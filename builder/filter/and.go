package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type And struct {
	Base
	Fields []builder.Filter `json:"fields,omitempty"`
}

func NewAnd() *And {
	a := &And{}
	a.SetType("and")
	return a
}

func (a *And) SetFields(fields []builder.Filter) *And {
	a.Fields = fields
	return a
}

func (a *And) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Fields []json.RawMessage `json:"fields,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var f builder.Filter
	ff := make([]builder.Filter, len(tmp.Fields))
	for i := range tmp.Fields {
		if f, err = Load(tmp.Fields[i]); err != nil {
			return err
		}
		ff[i] = f
	}
	a.Base = tmp.Base
	a.Fields = ff
	return nil
}
