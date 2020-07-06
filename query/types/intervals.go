package types

import (
	"time"
)

const (
	IntervalFormat = time.RFC3339
)

// Interval represents a druid interval.
type Interval struct {
	StartTime time.Time
	EndTime   time.Time
}

// NewInterval instantiate a new interval.
func NewInterval(startTime, endTime time.Time) Interval {
	return Interval{StartTime: startTime, EndTime: endTime}
}

// NewInterval is a helper to get slice of intervals.
func NewIntervals(ii ...Interval) []Interval {
	var intervals = make([]Interval, len(ii))
	for idx, i := range ii {
		intervals[idx] = i
	}
	return intervals
}

//MarshalJSON marshals Interval following ISO 8601 time interval.
func (i *Interval) MarshalText() ([]byte, error) {
	return []byte(i.StartTime.Format(IntervalFormat) + "/" + i.EndTime.Format(IntervalFormat)), nil
}
