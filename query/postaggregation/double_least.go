package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type DoubleLeast struct {
	Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewDoubleLeast() *DoubleLeast {
	d := &DoubleLeast{}
	d.SetType("doubleLeast")
	return d
}

func (d *DoubleLeast) SetName(name string) *DoubleLeast {
	d.Base.SetName(name)
	return d
}

func (d *DoubleLeast) SetFields(fields []query.PostAggregator) *DoubleLeast {
	d.Fields = fields
	return d
}

func (d *DoubleLeast) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Fields []json.RawMessage `json:"fields"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var p query.PostAggregator
	pp := make([]query.PostAggregator, len(tmp.Fields))
	for i := range tmp.Fields {
		if p, err = Load(tmp.Fields[i]); err != nil {
			return err
		}
		pp[i] = p
	}
	d.Base = tmp.Base
	d.Fields = pp
	return nil
}
