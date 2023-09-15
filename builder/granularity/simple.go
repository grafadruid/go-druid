package granularity

import "github.com/h2oai/go-druid/builder"

// Simple granularities are specified as a string and bucket timestamps by their UTC time.
// https://druid.apache.org/docs/latest/querying/granularities.html#simple-granularities
type Simple string

const (
	All           Simple = "all"
	None                 = "none"
	Second               = "second"
	Minute               = "minute"
	FifteenMinute        = "fifteen_minute"
	ThirtyMinute         = "thirty_minute"
	Hour                 = "hour"
	Day                  = "day"
	Week                 = "week"
	Month                = "month"
	Quarter              = "quarter"
	Year                 = "year"
)

// Type sets the type to Simple
func (s *Simple) Type() builder.ComponentType {
	return "simple"
}

// SetGranularity sets granularity.
func (s *Simple) SetGranularity(g Simple) *Simple {
	*s = g
	return s
}

// NewSimple creates a Simple type.
func NewSimple() *Simple {
	var s Simple
	return &s
}
