package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
)

// Extrema aggregations using generics for greatest/least operations

type NumericExtrema[T any] struct {
	Base
	Fields []builder.PostAggregator `json:"fields,omitempty"`
}

func (n *NumericExtrema[T]) SetName(name string) *NumericExtrema[T] {
	n.Base.SetName(name)
	return n
}

func (n *NumericExtrema[T]) SetFields(fields []builder.PostAggregator) *NumericExtrema[T] {
	n.Fields = fields
	return n
}

func (n *NumericExtrema[T]) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Fields []json.RawMessage `json:"fields,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var p builder.PostAggregator
	pp := make([]builder.PostAggregator, len(tmp.Fields))
	for i := range tmp.Fields {
		if p, err = Load(tmp.Fields[i]); err != nil {
			return err
		}
		pp[i] = p
	}
	n.Base = tmp.Base
	n.Fields = pp
	return nil
}

// Type aliases for backward compatibility
type (
	DoubleGreatest = NumericExtrema[float64]
	DoubleLeast    = NumericExtrema[float64]
	LongGreatest   = NumericExtrema[int64]
	LongLeast      = NumericExtrema[int64]
)

// Constructor functions
func NewDoubleGreatest() *DoubleGreatest {
	d := &DoubleGreatest{}
	d.SetType("doubleGreatest")
	return d
}

func NewDoubleLeast() *DoubleLeast {
	d := &DoubleLeast{}
	d.SetType("doubleLeast")
	return d
}

func NewLongGreatest() *LongGreatest {
	l := &LongGreatest{}
	l.SetType("longGreatest")
	return l
}

func NewLongLeast() *LongLeast {
	l := &LongLeast{}
	l.SetType("longLeast")
	return l
}
