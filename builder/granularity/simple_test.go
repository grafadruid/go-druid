package granularity

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewSimple(t *testing.T) {
	s := NewSimple()
	s.SetGranularity("all")

	expected := `"all"`
	var unmarshalled *Simple
	err := json.Unmarshal([]byte(expected), &unmarshalled)
	if err != nil {
		t.Errorf("json.Marshal failed, %s", err)
	}

	if !reflect.DeepEqual(s, unmarshalled) {
		generated, err := json.Marshal(s)
		t.Errorf("Expected=%s, Got=%s (err:%v)", expected, generated, err)
	}

}
