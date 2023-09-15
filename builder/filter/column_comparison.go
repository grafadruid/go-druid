package filter

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/dimension"
)

type ColumnComparison struct {
	Base
	Dimensions []builder.Dimension `json:"dimensions,omitempty"`
}

func NewColumnComparison() *ColumnComparison {
	c := &ColumnComparison{}
	c.SetType("columnComparison")
	return c
}

func (c *ColumnComparison) SetDimensions(dimensions []builder.Dimension) *ColumnComparison {
	c.Dimensions = dimensions
	return c
}

func (c *ColumnComparison) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		Base
		Dimensions []json.RawMessage `json:"dimensions,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var d builder.Dimension
	dd := make([]builder.Dimension, len(tmp.Dimensions))
	for i := range tmp.Dimensions {
		if d, err = dimension.Load(tmp.Dimensions[i]); err != nil {
			return err
		}
		dd[i] = d
	}
	c.Base = tmp.Base
	c.Dimensions = dd
	return nil
}
