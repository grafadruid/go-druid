package granularity

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNewPeriod(t *testing.T) {
	p := NewPeriod()
	p.SetOrigin(time.Unix(1613779200, 0))
	p.SetTimeZone(`America/Chicago`)
	p.SetPeriod(time.Minute)

	expected := `{"type":"period","period":60000000000,"origin":"2021-02-19T16:00:00-08:00","timeZone":"America/Chicago"}`
	var unmarshalled *Period
	err := json.Unmarshal([]byte(expected), &unmarshalled)
	if err != nil {
		t.Errorf("json.Unmarshal failed, %s", err)
	}

	if !reflect.DeepEqual(p, unmarshalled) {
		generated, err := json.Marshal(p)
		t.Errorf("Expected=%s, Got=%s (err:%v)", expected, generated, err)
	}
}
