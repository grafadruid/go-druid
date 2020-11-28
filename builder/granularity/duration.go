package granularity

import (
	"time"
)

type Duration struct {
	Base
	Duration time.Duration `json:"duration,omitempty"`
	Origin   time.Time     `json:"origin,omitempty"`
}

func NewDuration() *Duration {
	d := &Duration{}
	d.SetType("duration")
	return d
}

func (d *Duration) SetDuration(duration time.Duration) *Duration {
	d.Duration = duration
	return d
}

func (d *Duration) SetOrigin(origin time.Time) *Duration {
	d.Origin = origin
	return d
}
