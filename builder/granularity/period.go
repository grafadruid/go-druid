package granularity

import (
	"time"

	"github.com/h2oai/go-druid/builder/types"
)

// Period granularity is specified as arbitrary period combinations of years, months, weeks, hours, minutes and seconds
// (e.g. P2W, P3M, PT1H30M, PT0.750S) in ISO8601 format.
// https://druid.apache.org/docs/latest/querying/granularities.html#period-granularities
type Period struct {
	Base
	Period   time.Duration      `json:"period,omitempty"`
	Origin   time.Time          `json:"origin,omitempty"`
	TimeZone types.DateTimeZone `json:"timeZone,omitempty"`
}

// NewPeriod creates a Period type.
func NewPeriod() *Period {
	p := &Period{}
	p.SetType("period")
	return p
}

// SetPeriod sets period.
func (p *Period) SetPeriod(period time.Duration) *Period {
	p.Period = period
	return p
}

// SetOrigin sets origin.
func (p *Period) SetOrigin(origin time.Time) *Period {
	p.Origin = origin
	return p
}

// SetTimeZone sets timezone.
func (p *Period) SetTimeZone(timeZone types.DateTimeZone) *Period {
	p.TimeZone = timeZone
	return p
}
