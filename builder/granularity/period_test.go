package granularity

import (
	"github.com/grafadruid/go-druid/builder/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPeriod(t *testing.T) {
	p := NewPeriod()
	p.SetOrigin(time.Unix(1613779200, 0))
	p.SetTimeZone(`America/Chicago`)
	p.SetPeriod(time.Minute)

	expected := []byte(`{"type":"period","period":60000000000,"origin":"2021-02-19T16:00:00-08:00","timeZone":"America/Chicago"}`)

	built, err := Load(expected)
	assert.Nil(t, err)

	testutil.Compare(t, expected, p, built)
}
