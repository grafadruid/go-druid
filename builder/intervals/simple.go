package intervals

import "github.com/grafadruid/go-druid/builder"

type Simple []*Interval

func NewSimple() *Simple {
	s := &Simple{}
	return s
}

func (s *Simple) Type() builder.ComponentType {
	return "simple"
}

func (s *Simple) SetIntervals(intervals []*Interval) *Simple {
	*s = intervals
	return s
}
