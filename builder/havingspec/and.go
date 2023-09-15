package havingspec

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type And struct {
	Base
	HavingSpecs []builder.HavingSpec `json:"havingSpecs,omitempty"`
}

func NewAnd() *And {
	a := &And{}
	a.SetType("and")
	return a
}

func (a *And) SetHavingSpecs(havingSpecs []builder.HavingSpec) *And {
	a.HavingSpecs = havingSpecs
	return a
}

func (a *And) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		HavingSpecs []json.RawMessage `json:"havingSpecs,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var h builder.HavingSpec
	hh := make([]builder.HavingSpec, len(tmp.HavingSpecs))
	for i := range tmp.HavingSpecs {
		if h, err = Load(tmp.HavingSpecs[i]); err != nil {
			return err
		}
		hh[i] = h
	}
	a.Base = tmp.Base
	a.HavingSpecs = hh
	return nil
}
