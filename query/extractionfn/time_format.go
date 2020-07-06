package extractionfn

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type TimeFormat struct {
	*Base
	Format      string             `json:"format"`
	TimeZone    types.DateTimeZone `json:"timeZone"`
	Locale      string             `json:"string"`
	Granularity query.Granularity  `json:"granularity"`
	AsMillis    bool               `json:"asMillis"`
}

func NewTimeFormat() *TimeFormat {
	t := &TimeFormat{}
	t.SetType("timeFormat")
	return t
}

func (t *TimeFormat) SetFormat(format string) *TimeFormat {
	t.Format = format
	return t
}

func (t *TimeFormat) SetTimeZone(timeZone types.DateTimeZone) *TimeFormat {
	t.TimeZone = timeZone
	return t
}

func (t *TimeFormat) SetLocale(locale string) *TimeFormat {
	t.Locale = locale
	return t
}

func (t *TimeFormat) SetGranularity(granularity query.Granularity) *TimeFormat {
	t.Granularity = granularity
	return t
}

func (t *TimeFormat) SetAsMillis(asMillis bool) *TimeFormat {
	t.AsMillis = asMillis
	return t
}
