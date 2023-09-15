package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type Or struct {
	Base
	Fields []builder.Filter `json:"fields,omitempty"`
}

func NewOr() *Or {
	o := &Or{}
	o.SetType("or")
	return o
}

func (o *Or) SetFields(fields []builder.Filter) *Or {
	o.Fields = fields
	return o
}

func (o *Or) UnmarshalJSON(data []byte) error {
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
	o.Base = tmp.Base
	o.Fields = ff
	return nil
}
