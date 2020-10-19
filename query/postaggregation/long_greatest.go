package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type LongGreatest struct {
	Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewLongGreatest() *LongGreatest {
	l := &LongGreatest{}
	l.SetType("longGreatest")
	return l
}

func (l *LongGreatest) SetName(name string) *LongGreatest {
	l.Base.SetName(name)
	return l
}

func (l *LongGreatest) SetFields(fields []query.PostAggregator) *LongGreatest {
	l.Fields = fields
	return l
}

func (l *LongGreatest) UnmarshalJSON(data []byte) error {
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
	l.Base = tmp.Base
	l.Fields = pp
	return nil
}
