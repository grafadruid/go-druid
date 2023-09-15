package extractionfn

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/types"
)

type TimeFormat struct {
	Base
	Format      string              `json:"format,omitempty"`
	TimeZone    types.DateTimeZone  `json:"timeZone,omitempty"`
	Locale      string              `json:"locale,omitempty"`
	Granularity builder.Granularity `json:"granularity,omitempty"`
	AsMillis    *bool               `json:"asMillis,omitempty"`
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
	t.AsMillis = &asMillis
	return t
}

func (t *TimeFormat) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Format      string             `json:"format,omitempty"`
		TimeZone    types.DateTimeZone `json:"timeZone,omitempty"`
		Locale      string             `json:"locale,omitempty"`
		Granularity json.RawMessage    `json:"granularity,omitempty"`
		AsMillis    *bool              `json:"asMillis,omitempty"`
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
