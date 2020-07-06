package extractionfn

import "github.com/grafadruid/go-druid/query"

type Lookup struct {
	*Base
	Lookup                  query.LookupExtractor `json:"lookup"`
	RetainMissingValue      bool                  `json:"retainMissingValue"`
	ReplaceMissingValueWith string                `json:"replaceMissingValueWith,omitempty"`
	Injective               bool                  `json:"injective,omitempty"`
	Optimize                bool                  `json:"optimize,omitempty"`
}

func NewLookup() *Lookup {
	l := &Lookup{}
	l.SetType("lookup")
	return l
}

func (l *Lookup) SetLookup(lookup query.LookupExtractor) *Lookup {
	l.Lookup = lookup
	return l
}

func (l *Lookup) SetRetainMissingValue(retainMissingValue bool) *Lookup {
	l.RetainMissingValue = retainMissingValue
	return l
}

func (l *Lookup) SetReplaceMissingValueWith(replaceMissingValueWith string) *Lookup {
	l.ReplaceMissingValueWith = replaceMissingValueWith
	return l
}

func (l *Lookup) SetInjective(injective bool) *Lookup {
	l.Injective = injective
	return l
}

func (l *Lookup) SetOptimize(optimize bool) *Lookup {
	l.Optimize = optimize
	return l
}
