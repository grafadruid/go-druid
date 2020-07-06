package aggregation

type Filtered struct {
	*Base
	Aggregator string
	Filter     string
}

func NewFiltered() *Filtered {
	f := &Filtered{}
	f.SetType("filtered")
	return f
}

func (f *Filtered) SetName(name string) *Filtered {
	f.Base.SetName(name)
	return f
}
