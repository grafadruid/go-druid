package extractionfn

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/lookup"
)

type Lookup struct {
	Base
	Lookup                  builder.LookupExtractor `json:"lookup,omitempty"`
	RetainMissingValue      *bool                   `json:"retainMissingValue,omitempty"`
	ReplaceMissingValueWith string                  `json:"replaceMissingValueWith,omitempty"`
	Injective               *bool                   `json:"injective,omitempty"`
	Optimize                *bool                   `json:"optimize,omitempty"`
}

func NewLookup() *Lookup {
	l := &Lookup{}
	l.SetType("lookup")
	return l
}

func (l *Lookup) SetLookup(lookup builder.LookupExtractor) *Lookup {
	l.Lookup = lookup
	return l
}

func (l *Lookup) SetRetainMissingValue(retainMissingValue bool) *Lookup {
	l.RetainMissingValue = &retainMissingValue
	return l
}

func (l *Lookup) SetReplaceMissingValueWith(replaceMissingValueWith string) *Lookup {
	l.ReplaceMissingValueWith = replaceMissingValueWith
	return l
}

func (l *Lookup) SetInjective(injective bool) *Lookup {
	l.Injective = &injective
	return l
}

func (l *Lookup) SetOptimize(optimize bool) *Lookup {
	l.Optimize = &optimize
	return l
}
func (l *Lookup) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Lookup                  json.RawMessage `json:"lookup,omitempty"`
		RetainMissingValue      *bool           `json:"retainMissingValue,omitempty"`
		ReplaceMissingValueWith string          `json:"replaceMissingValueWith,omitempty"`
		Injective               *bool           `json:"injective,omitempty"`
		Optimize                *bool           `json:"optimize,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	lu, err := lookup.Load(tmp.Lookup)
	if err != nil {
		return err
	}
	l.Base = tmp.Base
	l.Lookup = lu
	l.RetainMissingValue = tmp.RetainMissingValue
	l.ReplaceMissingValueWith = tmp.ReplaceMissingValueWith
	l.Injective = tmp.Injective
	l.Optimize = tmp.Optimize
	return nil
}
