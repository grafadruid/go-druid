package havingspec

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/extractionfn"
)

type DimSelector struct {
	Base
	Dimension    string               `json:"dimension,omitempty"`
	Value        string               `json:"value,omitempty"`
	ExtractionFn builder.ExtractionFn `json:"extractionFn,omitempty"`
}

func NewDimSelector() *DimSelector {
	d := &DimSelector{}
	d.SetType("dimSelector")
	return d
}

func (d *DimSelector) SetDimension(dimension string) *DimSelector {
	d.Dimension = dimension
	return d
}

func (d *DimSelector) SetValue(value string) *DimSelector {
	d.Value = value
	return d
}

func (d *DimSelector) SetExtractionFn(extractionFn builder.ExtractionFn) *DimSelector {
	d.ExtractionFn = extractionFn
	return d
}

func (d *DimSelector) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension,omitempty"`
		Value        string          `json:"value,omitempty"`
		ExtractionFn json.RawMessage `json:"extractionFn,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	e, err := extractionfn.Load(tmp.ExtractionFn)
	if err != nil {
		return err
	}
	d.Base = tmp.Base
	d.Dimension = tmp.Dimension
	d.Value = tmp.Value
	d.ExtractionFn = e
	return nil
}
