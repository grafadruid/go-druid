package intervals

import (
	"strings"
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
func NewInterval(startTime, endTime time.Time) *Interval {
	return &Interval{StartTime: startTime, EndTime: endTime}
}

// MarshalText marshals Interval following ISO 8601 time interval.
func (i *Interval) MarshalText() ([]byte, error) {
	return []byte(i.StartTime.Format(IntervalFormat) + "/" + i.EndTime.Format(IntervalFormat)), nil
}

func (i *Interval) UnmarshalText(text []byte) error {
	interval := strings.Split(string(text), "/")
	var err error
	i.StartTime, err = time.Parse(time.RFC3339Nano, interval[0])
	if err != nil {
		return err
	}
	i.EndTime, err = time.Parse(time.RFC3339Nano, interval[1])
	if err != nil {
		return err
	}
	return nil
}
