package granularity

import (
	"time"
)

// Duration granularity is specified as an exact duration in milliseconds and timestamps are returned as UTC.
// Duration granularity values are in millis.
// https://druid.apache.org/docs/latest/querying/granularities.html#duration-granularities
type Duration struct {
	Base
	Duration time.Duration `json:"duration,omitempty"`
	Origin   time.Time     `json:"origin,omitempty"`
}

// NewDuration creates new Duration.
func NewDuration() *Duration {
	d := &Duration{}
	d.SetType("duration")
	return d
}

// SetDuration sets duration.
func (d *Duration) SetDuration(duration time.Duration) *Duration {
	d.Duration = duration
	return d
}

// SetOrigin sets the origin
func (d *Duration) SetOrigin(origin time.Time) *Duration {
	d.Origin = origin
	return d
}
