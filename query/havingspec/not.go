package havingspec

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Not struct {
	Base
	HavingSpec query.HavingSpec `json:"havingSpec"`
}

func NewNot() *Not {
	n := &Not{}
	n.SetType("not")
	return n
}

func (n *Not) SetHavingSpecs(havingSpec query.HavingSpec) *Not {
	n.HavingSpec = havingSpec
	return n
}

func (n *Not) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		HavingSpec json.RawMessage `json:"havingSpec"`
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
