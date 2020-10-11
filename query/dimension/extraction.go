package dimension

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Extraction struct {
	Base
	ExtractionFn query.ExtractionFn `json:"extractionFn"`
}

func NewExtraction() *Extraction {
	e := &Extraction{}
	e.SetType("extraction")
	return e
}

func (e *Extraction) SetDimension(dimension string) *Extraction {
	e.Base.SetDimension(dimension)
	return e
}

func (e *Extraction) SetOutputName(outputName string) *Extraction {
	e.Base.SetOutputName(outputName)
	return e
}

func (e *Extraction) SetOutputType(outputType types.OutputType) *Extraction {
	e.Base.SetOutputType(outputType)
	return e
}

func (e *Extraction) SetExtractionFn(extractionFn query.ExtractionFn) *Extraction {
	e.ExtractionFn = extractionFn
	return e
}
