package granularity

import "github.com/grafadruid/go-druid/builder"

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

func (s *Simple) Type() builder.ComponentType {
	return "simple"
}

func (s *Simple) SetGranularity(g Simple) *Simple {
	*s = g
	return s
}

func NewSimple() *Simple {
	var s Simple
	return &s
}

// TypeComplexSimple is for those Simple types than can take both Long and Short specs
type TypeComplexSimple struct {
	Base
}

func NewComplexSimple() *TypeComplexSimple {
	var a TypeComplexSimple
	return &a
}

func (a *TypeComplexSimple) Type() builder.ComponentType {
	return "all"
}
