package postaggregation

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type DoubleLeast struct {
	Base
	Fields []builder.PostAggregator `json:"fields,omitempty"`
}

func NewDoubleLeast() *DoubleLeast {
	d := &DoubleLeast{}
	d.SetType("doubleLeast")
	return d
}

func (d *DoubleLeast) SetName(name string) *DoubleLeast {
	d.Base.SetName(name)
	return d
}

func (d *DoubleLeast) SetFields(fields []builder.PostAggregator) *DoubleLeast {
	d.Fields = fields
	return d
}

func (d *DoubleLeast) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Fields []json.RawMessage `json:"fields,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var p builder.PostAggregator
	pp := make([]builder.PostAggregator, len(tmp.Fields))
	for i := range tmp.Fields {
		if p, err = Load(tmp.Fields[i]); err != nil {
			return err
		}
		pp[i] = p
	}
	d.Base = tmp.Base
	d.Fields = pp
	return nil
}
