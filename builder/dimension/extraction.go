package dimension

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/types"
)

type Extraction struct {
	Base
	ExtractionFn builder.ExtractionFn `json:"extractionFn,omitempty"`
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

func (e *Extraction) SetExtractionFn(extractionFn builder.ExtractionFn) *Extraction {
	e.ExtractionFn = extractionFn
	return e
}

func (e *Extraction) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		ExtractionFn json.RawMessage `json:"extractionFn,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	ef, err := Load(tmp.ExtractionFn)
	if err != nil {
		return err
	}
	e.Base = tmp.Base
	e.ExtractionFn = ef
	return nil
}
