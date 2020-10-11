package granularity

import (
	"time"

	"github.com/grafadruid/go-druid/query/types"
)

type Period struct {
	Base
	Period   time.Duration      `json:"period"`
	Origin   time.Time          `json:"origin"`
	TimeZone types.DateTimeZone `json:"timeZone"`
}

func NewPeriod() *Period {
	p := &Period{}
	p.SetType("period")
	return p
}

func (p *Period) SetPeriod(period time.Duration) *Period {
	p.Period = period
	return p
}

func (p *Period) SetOrigin(origin time.Time) *Period {
	p.Origin = origin
	return p
}

func (p *Period) SetTimeZone(timeZone types.DateTimeZone) *Period {
	p.TimeZone = timeZone
	return p
}
