package extractionfn

import "github.com/grafadruid/go-druid/query"

type Cascade struct {
	Base
	ExtractionFns []query.ExtractionFn `json:"extractionFns"`
}

func NewCascade() *Cascade {
	c := &Cascade{}
	c.SetType("cascade")
	return c
}

func (c *Cascade) SetExtractionFns(extractionFns []query.ExtractionFn) *Cascade {
	c.ExtractionFns = extractionFns
	return c
}
