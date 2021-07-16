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

	x := &Duration{
		Base:     Base{Typ: "duration"},
		Duration: 60000000000,
		Origin:   time.Unix(1613779200, 0),
	}
	assert.Equal(t, x, d)
}
