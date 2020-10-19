package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type LongLeast struct {
	Base
	Fields []query.PostAggregator `json:"fields"`
}

func NewLongLeast() *LongLeast {
	l := &LongLeast{}
	l.SetType("longLeast")
	return l
}

func (l *LongLeast) SetName(name string) *LongLeast {
	l.Base.SetName(name)
	return l
}

func (l *LongLeast) SetFields(fields []query.PostAggregator) *LongLeast {
	l.Fields = fields
	return l
}

func (l *LongLeast) UnmarshalJSON(data []byte) error {
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
