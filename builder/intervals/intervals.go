package intervals

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type,omitempty"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.Intervals, error) {
	var i builder.Intervals
	if string(data) == "null" {
		return i, nil
	}
	// "intervals" in the spec is just an string array
	var intv []string
	if err := json.Unmarshal(data, &intv); err != nil {
		return nil, err
	}
	// Now cast the only array item into an actual "interval"
	interval := Interval(intv[0])
	// create our "intervals" object with "Typ"
	i = &Intervals{
		Base: Base{
			Typ: "intervals",
		},
		Intervals: []*Interval{&interval},
	}
	return i, nil
}
