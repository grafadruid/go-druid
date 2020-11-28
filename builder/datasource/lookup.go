package datasource

type Lookup struct {
	Base
	Lookup string `json:"lookup,omitempty"`
}

func NewLookup() *Lookup {
	l := &Lookup{}
	l.SetType("lookup")
	return l
}

func (l *Lookup) SetLookup(lookup string) *Lookup {
	l.Lookup = lookup
	return l
}
