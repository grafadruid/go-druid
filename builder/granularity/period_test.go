package granularity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPeriod(t *testing.T) {
	p := NewPeriod()
	p.SetOrigin(time.Unix(1613779200, 0))
	p.SetTimeZone(`America/Chicago`)
	p.SetPeriod(time.Minute)

	x := &Period{
		Base:     Base{Typ: "period"},
		Period:   60000000000,
		Origin:   time.Unix(1613779200, 0),
		TimeZone: "America/Chicago",
	}
	assert.Equal(t, x, p)
}
