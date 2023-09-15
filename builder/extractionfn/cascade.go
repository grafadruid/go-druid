package extractionfn

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type Cascade struct {
	Base
	ExtractionFns []builder.ExtractionFn `json:"extractionFns,omitempty"`
}

func NewCascade() *Cascade {
	c := &Cascade{}
	c.SetType("cascade")
	return c
}

func (c *Cascade) SetExtractionFns(extractionFns []builder.ExtractionFn) *Cascade {
	c.ExtractionFns = extractionFns
	return c
}

func (c *Cascade) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		ExtractionFns []json.RawMessage `json:"extractionFns,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var e builder.ExtractionFn
	ee := make([]builder.ExtractionFn, len(tmp.ExtractionFns))
	for i := range tmp.ExtractionFns {
		if e, err = Load(tmp.ExtractionFns[i]); err != nil {
			return err
		}
		ee[i] = e
	}
	c.Base = tmp.Base
	c.ExtractionFns = ee
	return nil
}
