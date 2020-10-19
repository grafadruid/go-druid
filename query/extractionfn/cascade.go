package extractionfn

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

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

func (c *Cascade) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		ExtractionFns []json.RawMessage `json:"extractionFns"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var e query.ExtractionFn
	ee := make([]query.ExtractionFn, len(tmp.ExtractionFns))
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
