package havingspec

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Or struct {
	Base
	HavingSpecs []query.HavingSpec `json:"havingSpecs"`
}

func NewOr() *Or {
	o := &Or{}
	o.SetType("or")
	return o
}

func (o *Or) SetHavingSpecs(havingSpecs []query.HavingSpec) *Or {
	o.HavingSpecs = havingSpecs
	return o
}

func (o *Or) UnmarshalJSON(data []byte) error {
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
	o.Base = tmp.Base
	o.HavingSpecs = hh
	return nil
}
