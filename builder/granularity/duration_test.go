package granularity

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder"
	"reflect"
	"testing"
	"time"
)

func TestNewDuration(t *testing.T) {
	d := NewDuration()
	d.SetDuration(time.Minute)
	d.SetOrigin(time.Unix(1613779200, 0))

	expected := `{"type":"duration","duration":60000000000,"origin":"2021-02-19T16:00:00-08:00"}`
	var unmarshalled builder.Granularity
	unmarshalled, err := Load([]byte(expected))
	if err != nil {
		t.Errorf("Load failed, %s", err)
	}

	if !reflect.DeepEqual(d, unmarshalled) {
		generated, err := json.Marshal(d)
		t.Errorf("Expected=%s, Got=%s (err:%v)", expected, generated, err)
	}
}
