package aggregation

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/filter"
)

type Filtered struct {
	Base
	Aggregator builder.Aggregator `json:"aggregator,omitempty"`
	Filter     builder.Filter     `json:"filter,omitempty"`
}

func NewFiltered() *Filtered {
	f := &Filtered{}
	f.SetType("filtered")
	return f
}

func (f *Filtered) SetName(name string) *Filtered {
	f.Base.SetName(name)
	return f
}

func (f *Filtered) SetAggregator(aggregator builder.Aggregator) *Filtered {
	f.Aggregator = aggregator
	return f
}

func (f *Filtered) SetFilter(filter builder.Filter) *Filtered {
	f.Filter = filter
	return f
}

func (f *Filtered) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Aggregator json.RawMessage `json:"aggregator,omitempty"`
		Filter     json.RawMessage `json:"filter,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	a, err := Load(tmp.Aggregator)
	if err != nil {
		return err
	}
	filter, err := filter.Load(tmp.Filter)
	if err != nil {
		return err
	}
	f.Base = tmp.Base
	f.Aggregator = a
	f.Filter = filter
	return nil
}
