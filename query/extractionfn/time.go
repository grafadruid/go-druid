package extractionfn

import "github.com/grafadruid/go-druid/query/types"

type Time struct {
	*Base
	TimeFormat   string             `json:"timeFormat"`
	ResultFormat types.ResultFormat `json:"resultFormat"`
	Joda         bool               `json:"joda"`
}

func NewTime() *Time {
	t := &Time{}
	t.SetType("time")
	return t
}

func (t *Time) SetTimeFormat(timeFormat string) *Time {
	t.TimeFormat = timeFormat
	return t
}

func (t *Time) SetResultFormat(resultFormat types.ResultFormat) *Time {
	t.ResultFormat = resultFormat
	return t
}

func (t *Time) SetJoda(joda bool) *Time {
	t.Joda = joda
	return t
}
