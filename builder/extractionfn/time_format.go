package extractionfn

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/types"
)

type TimeFormat struct {
	Base
	Format      string              `json:"format"`
	TimeZone    types.DateTimeZone  `json:"timeZone"`
	Locale      string              `json:"locale"`
	Granularity builder.Granularity `json:"granularity"`
	AsMillis    bool                `json:"asMillis"`
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

func (t *TimeFormat) SetGranularity(granularity builder.Granularity) *TimeFormat {
	t.Granularity = granularity
	return t
}

func (t *TimeFormat) SetAsMillis(asMillis bool) *TimeFormat {
	t.AsMillis = asMillis
	return t
}

func (t *TimeFormat) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Format      string             `json:"format"`
		TimeZone    types.DateTimeZone `json:"timeZone"`
		Locale      string             `json:"locale"`
		Granularity json.RawMessage    `json:"granularity"`
		AsMillis    bool               `json:"asMillis"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	g, err := Load(tmp.Granularity)
	if err != nil {
		return err
	}
	t.Base = tmp.Base
	t.Format = tmp.Format
	t.TimeZone = tmp.TimeZone
	t.Locale = tmp.Locale
	t.Granularity = g
	t.AsMillis = tmp.AsMillis
	return nil
}
