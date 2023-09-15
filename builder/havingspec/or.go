package havingspec

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type Or struct {
	Base
	HavingSpecs []builder.HavingSpec `json:"havingSpecs,omitempty"`
}

func NewOr() *Or {
	o := &Or{}
	o.SetType("or")
	return o
}

func (o *Or) SetHavingSpecs(havingSpecs []builder.HavingSpec) *Or {
	o.HavingSpecs = havingSpecs
	return o
}

func (o *Or) UnmarshalJSON(data []byte) error {
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
	o.Base = tmp.Base
	o.HavingSpecs = hh
	return nil
}
