package extractionfn

// RegisteredLookup holds the registered lookup extraction function struct based on
// https://druid.apache.org/docs/latest/querying/dimensionspecs.html#registered-lookup-extraction-function
type RegisteredLookup struct {
	Base
	Lookup                  string `json:"lookup,omitempty"`
	RetainMissingValue      *bool  `json:"retainMissingValue,omitempty"`
	ReplaceMissingValueWith string `json:"replaceMissingValueWith,omitempty"`
	Injective               *bool  `json:"injective,omitempty"`
	Optimize                *bool  `json:"optimize,omitempty"`
}

func NewRegisteredLookup() *RegisteredLookup {
	l := &RegisteredLookup{}
	l.SetType("registeredLookup")
	return l
}

func (l *RegisteredLookup) SetLookup(lookup string) *RegisteredLookup {
	l.Lookup = lookup
	return l
}

func (l *RegisteredLookup) SetRetainMissingValue(retainMissingValue bool) *RegisteredLookup {
	l.RetainMissingValue = &retainMissingValue
	return l
}

func (l *RegisteredLookup) SetReplaceMissingValueWith(replaceMissingValueWith string) *RegisteredLookup {
	l.ReplaceMissingValueWith = replaceMissingValueWith
	return l
}

func (l *RegisteredLookup) SetInjective(injective bool) *RegisteredLookup {
	l.Injective = &injective
	return l
}

func (l *RegisteredLookup) SetOptimize(optimize bool) *RegisteredLookup {
	l.Optimize = &optimize
	return l
}
