package aggregation

type Histogram struct {
	Base
	FieldName string    `json:"fieldName,omitempty"`
	Breaks    []float64 `json:"breaks,omitempty"`
}

func NewHistogram() *Histogram {
	h := &Histogram{}
	h.SetType("histogram")
	return h
}

func (h *Histogram) SetName(name string) *Histogram {
	h.Base.SetName(name)
	return h
}

func (h *Histogram) SetFieldName(fieldName string) *Histogram {
	h.FieldName = fieldName
	return h
}

func (h *Histogram) SetBreaks(breaks []float64) *Histogram {
	h.Breaks = breaks
	return h
}
