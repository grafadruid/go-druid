package filter

import (
	"encoding/json"
	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/intervals"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f,
		"filter should be nil")
	assert.NotNil(err,
		"error should not be nil")
	assert.Error(err,
		"unsupported filter type")
}

func TestNewInterval(t *testing.T) {
	location, _ := time.LoadLocation("UTC")
	start, _ := time.ParseInLocation(time.RFC3339Nano,
		"2022-06-16T08:28:53.33441Z",
		location)
	end, _ := time.ParseInLocation(time.RFC3339Nano,
		"2022-06-16T15:28:53.33441Z",
		location)
	// simple interval
	i := intervals.NewInterval().SetInterval(start,
		end)
	filter1 := NewSelector().SetDimension("countryName").SetValue("France")
	filterInterval := NewInterval().SetIntervals([]*intervals.Interval{i}).SetDimension("__time")
	filters := NewOr().SetFields([]builder.Filter{filter1, filterInterval})

	t.Run("marshal filter with interval",
		func(t *testing.T) {
			f, err := json.Marshal(filters)
			assert.Nil(t,
				err)

			assert.Nil(t,
				err)
			assert.Equal(t,
				`{"type":"or","fields":[{"type":"selector","dimension":"countryName","value":"France"},{"type":"interval","dimension":"__time","intervals":["2022-06-16T08:28:53.33441Z/2022-06-16T15:28:53.33441Z"]}]}`,
				string(f),
				"filter with time interval")
		})

	t.Run("marshal load marshal filter with interval",
		func(t *testing.T) {
			f, err := json.Marshal(filters)
			assert.Nil(t,
				err)

			filterWithIntervalObj, err := Load(f)
			assert.Nil(t,
				err)
			assert.NotNil(t,
				filterWithIntervalObj)
			assert.Equal(t,
				filters,
				filterWithIntervalObj)

			fJson, err := json.Marshal(filterWithIntervalObj)
			assert.Nil(t,
				err)

			assert.Nil(t,
				err)
			assert.Equal(t,
				`{"type":"or","fields":[{"type":"selector","dimension":"countryName","value":"France"},{"type":"interval","dimension":"__time","intervals":["2022-06-16T08:28:53.33441Z/2022-06-16T15:28:53.33441Z"]}]}`,
				string(fJson),
				"filter with time interval")
		})
}
