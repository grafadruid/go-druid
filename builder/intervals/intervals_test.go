package intervals

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntervals_Load(t *testing.T) {
	assert := assert.New(t)
	t.Run("test unsupported type",
		func(t *testing.T) {
			f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

			assert.Nil(f,
				"filter should be nil")
			assert.NotNil(err,
				"error should not be nil")
			assert.Error(err,
				"unsupported intervals type")
		})

	t.Run("simple interval not supported",
		func(t *testing.T) {
			f, err := Load([]byte(`"2022-06-16T08:28:53.33441Z/2022-06-16T15:28:53.33441Z"`))

			assert.Nil(f,
				"filter should be nil")
			assert.NotNil(err,
				"error should not be nil")
			assert.Error(err,
				"unsupported intervals type")
		})

	t.Run("complex interval supported",
		func(t *testing.T) {
			f, err := Load([]byte(`{"type":"intervals","intervals":["2022-06-16T08:28:53.33441Z/2022-06-16T15:28:53.33441Z"]}`))
			location, _ := time.LoadLocation("UTC")
			start, _ := time.ParseInLocation(time.RFC3339Nano,
				"2022-06-16T08:28:53.33441Z",
				location)
			end, _ := time.ParseInLocation(time.RFC3339Nano,
				"2022-06-16T15:28:53.33441Z",
				location)
			// simple interval
			interval := NewInterval().SetInterval(start,
				end)
			// complex intervals
			intervals := NewIntervals().SetIntervals([]*Interval{interval})

			assert.Nil(err,
				"error should be nil")
			assert.Equal(intervals,
				f,
				"loaded intervals match the built intervals")
		})
}

func TestIntervals_MarshalJSON(t *testing.T) {
	location, _ := time.LoadLocation("UTC")
	start, _ := time.ParseInLocation(time.RFC3339Nano,
		"2022-06-16T08:28:53.33441Z",
		location)
	end, _ := time.ParseInLocation(time.RFC3339Nano,
		"2022-06-16T15:28:53.33441Z",
		location)
	// simple interval
	interval := NewInterval().SetInterval(start,
		end)
	// complex intervals
	intervals := NewIntervals().SetIntervals([]*Interval{interval})

	t.Run("simple interval returns a string",
		func(t *testing.T) {
			f, err := json.Marshal(interval)
			if err != nil {
				fmt.Println(err.Error())
			}
			assert.Nil(t,
				err)
			assert.Equal(t,
				`"2022-06-16T08:28:53.33441Z/2022-06-16T15:28:53.33441Z"`,
				string(f),
				"simple interval returns a string")
		})

	t.Run("marshal generates complex interval type",
		func(t *testing.T) {
			f, err := json.Marshal(intervals)
			if err != nil {
				fmt.Println(err.Error())
			}
			assert.Nil(t,
				err)
			assert.Equal(t,
				`{"type":"intervals","intervals":["2022-06-16T08:28:53.33441Z/2022-06-16T15:28:53.33441Z"]}`,
				string(f),
				"complex interval returns a struct")
		})
}
