package granularity

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

func NewSimple() *Simple {
	var s Simple
	return &s
}
