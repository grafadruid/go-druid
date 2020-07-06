package filter

import "github.com/grafadruid/go-druid/query"

type Extraction struct {
	*Base
	Dimension    string             `json:"dimension"`
	Value        string             `json:"value"`
	ExtractionFn query.ExtractionFn `json:"extractionFn"`
}

func NewExtraction() *Extraction {
	e := &Extraction{}
	e.SetType("extraction")
	return e
}

func (e *Extraction) SetDimension(dimension string) *Extraction {
	e.Dimension = dimension
	return e
}

func (e *Extraction) SetValue(value string) *Extraction {
	e.Value = value
	return e
}

func (e *Extraction) SetExtractionFn(extractionFn query.ExtractionFn) *Extraction {
	e.ExtractionFn = extractionFn
	return e
}
