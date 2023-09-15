package topnmetric

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
)

type Inverted struct {
	Base
	Metric builder.TopNMetric `json:"metric,omitempty"`
}

func NewInverted() *Inverted {
	i := &Inverted{}
	i.SetType("inverted")
	return i
}

func (i *Inverted) SetMetric(metric builder.TopNMetric) *Inverted {
	i.Metric = metric
	return i
}

func (i *Inverted) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Metric json.RawMessage `json:"metric,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	m, err := Load(tmp.Metric)
	if err != nil {
		return err
	}
	i.Base = tmp.Base
	i.Metric = m
	return nil
}
