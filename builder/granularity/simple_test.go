package granularity

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder"
	"reflect"
	"testing"
)

func TestNewSimple(t *testing.T) {
	s := NewSimple()
	s.SetGranularity("all")

	expected := `"all"`
	var unmarshalled builder.Granularity
	unmarshalled, err := Load([]byte(expected))
	if err != nil {
		t.Errorf("Load failed, %s", err)
	}

	if !reflect.DeepEqual(s, unmarshalled) {
		generated, err := json.Marshal(s)
		t.Errorf("Expected=%s, Got=%s (err:%v)", expected, generated, err)
	}

}
