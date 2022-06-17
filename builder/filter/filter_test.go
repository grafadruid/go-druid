package filter

import (
	"encoding/json"
	"fmt"
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/intervals"
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

	f, err := json.Marshal(filters)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t,
		err)
	assert.Equal(t,
		`{"type":"or","fields":[{"type":"selector","dimension":"countryName","value":"France"},{"type":"interval","dimension":"__time","intervals":["2022-06-16T08:28:53.33441Z/2022-06-16T15:28:53.33441Z"]}]}`,
		string(f),
		"filter with time interval")
}
