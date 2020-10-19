package dimension

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Lookup struct {
	Base
	Name                    string                `json:"name"`
	ReplaceMissingValueWith string                `json:"replaceMissingValueWith"`
	RetainMissingValue      bool                  `json:"retainMissingValue"`
	Lookup                  query.LookupExtractor `json:"lookup"`
	Optimize                bool                  `json:"optimize"`
}

type RegisteredLookup struct {
	Base
	Name string `json:"name"`
}

func NewLookup() *Lookup {
	l := &Lookup{}
	l.SetType("lookup")
	return l
}

func (l *Lookup) SetName(name string) *Lookup {
	l.Name = name
	return l
}

func (l *Lookup) SetReplaceMissingValueWith(replaceMissingValueWith string) *Lookup {
	l.ReplaceMissingValueWith = replaceMissingValueWith
	return l
}

func (l *Lookup) SetRetainMissingValue(retainMissingValue bool) *Lookup {
	l.RetainMissingValue = retainMissingValue
	return l
}

func (l *Lookup) SetLookup(lookup query.LookupExtractor) *Lookup {
	l.Lookup = lookup
	return l
}

func (l *Lookup) SetOptimize(optimize bool) *Lookup {
	l.Optimize = optimize
	return l
}

func (l *Lookup) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Name                    string          `json:"name"`
		ReplaceMissingValueWith string          `json:"replaceMissingValueWith"`
		RetainMissingValue      bool            `json:"retainMissingValue"`
		Lookup                  json.RawMessage `json:"lookup"`
		Optimize                bool            `json:"optimize"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	//l, err := lookupextractor.Load(tmp.Filter)
	//if err != nil {
	//	return err
	//}
	l.Base = tmp.Base
	l.Name = tmp.Name
	l.ReplaceMissingValueWith = tmp.ReplaceMissingValueWith
	l.RetainMissingValue = tmp.RetainMissingValue
	//l.Lookup = lookup
	l.Optimize = tmp.Optimize
	return nil
}
