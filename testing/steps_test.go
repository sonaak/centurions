package testing

import "testing"

func TestRecorder_Get(t *testing.T) {
	recorder := Recorder {
		store: make(map[string]interface{}),
	}

	recorder.Set("a", uint(1))
	val, err := recorder.Get("a")

	if err != nil {
		t.Errorf("expect no error when calling Get: %v", err)
	}

	switch v := val.(type) {
	case uint:
		if v != 1 {
			t.Errorf("expect value to be 1 (uint); got %d instead", v)
		}

	default:
		t.Errorf("expect value to be an integer")
	}
}

