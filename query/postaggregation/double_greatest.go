package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type DoubleGreatest struct {
	Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewDoubleGreatest() *DoubleGreatest {
	d := &DoubleGreatest{}
	d.SetType("doubleGreatest")
	return d
}

func (d *DoubleGreatest) SetName(name string) *DoubleGreatest {
	d.Base.SetName(name)
	return d
}

func (d *DoubleGreatest) SetFields(fields []query.PostAggregator) *DoubleGreatest {
	d.Fields = fields
	return d
}

func (d *DoubleGreatest) UnmarshalJSON(data []byte) error {
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
