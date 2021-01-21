package intervals

type Intervals struct {
	Base
	Intervals []*Interval `json:"intervals,omitempty"`
}

func NewIntervals() *Intervals {
	i := &Intervals{}
	i.SetType("intervals")
	return i
}

func (i *Intervals) SetIntervals(intervals []*Interval) *Intervals {
	i.Intervals = intervals
	return i
}
