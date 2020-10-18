package havingspec

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type And struct {
	Base
	HavingSpecs []query.HavingSpec `json:"havingSpecs"`
}

func NewAnd() *And {
	a := &And{}
	a.SetType("and")
	return a
}

func (a *And) SetHavingSpecs(havingSpecs []query.HavingSpec) *And {
	a.HavingSpecs = havingSpecs
	return a
}

func (a *And) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		HavingSpecs []json.RawMessage `json:"havingSpecs"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var h query.HavingSpec
	hh := make([]query.HavingSpec, len(tmp.HavingSpecs))
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
