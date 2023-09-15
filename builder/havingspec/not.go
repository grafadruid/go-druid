package havingspec

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type Not struct {
	Base
	HavingSpec builder.HavingSpec `json:"havingSpec,omitempty"`
}

func NewNot() *Not {
	n := &Not{}
	n.SetType("not")
	return n
}

func (n *Not) SetHavingSpecs(havingSpec builder.HavingSpec) *Not {
	n.HavingSpec = havingSpec
	return n
}

func (n *Not) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		HavingSpec json.RawMessage `json:"havingSpec,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	h, err := Load(tmp.HavingSpec)
	if err != nil {
		return err
	}
	n.Base = tmp.Base
	n.HavingSpec = h
	return nil
}
