package intervals

import (
	"time"
)

const (
	IntervalTimeFormat = time.RFC3339Nano
)

// Interval represents a druid interval.
type Interval string

// NewInterval instantiate a new interval.
func NewInterval() *Interval {
	var i Interval
	return &i
}

func (i *Interval) SetInterval(start, end time.Time) *Interval {
	*i = Interval(start.Format(IntervalTimeFormat) + "/" + end.Format(IntervalTimeFormat))
	return i
}

func (i *Interval) SetIntervalWithString(start, end string) *Interval {
	*i = Interval(start + "/" + end)
	return i
}
