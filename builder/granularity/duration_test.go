package granularity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDuration(t *testing.T) {
	d := NewDuration()
	d.SetDuration(time.Minute)
	d.SetOrigin(time.Unix(1613779200, 0))
	start, _ := time.Parse(time.RFC822, "19 Feb 21 19:00 EST")

	x := &Duration{
		Base:     Base{Typ: "duration"},
		Duration: 60000000000,
		Origin:   start,
	}
	assert.Equal(t, x, d)
}
