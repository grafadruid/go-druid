package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type Not struct {
	Base
	Field builder.Filter `json:"field,omitempty"`
}

func NewNot() *Not {
	n := &Not{}
	n.SetType("not")
	return n
}

func (n *Not) SetField(field builder.Filter) *Not {
	n.Field = field
	return n
}

func (n *Not) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Field json.RawMessage `json:"field,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var f builder.Filter
	if tmp.Field != nil {
		f, err = Load(tmp.Field)
		if err != nil {
			return err
		}
	}
	n.Base = tmp.Base
	n.Field = f
	return nil
}
