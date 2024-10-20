package maxvalue

import "testing"

func TestMaxValue(t *testing.T) {
	var expected float32 = 42.42
	got := MaxValue([]float32{41.2, 42.4, 35.2, 42.42, 4.0})
	if expected != got {
		t.Errorf("expected %f but got %f", expected, got)
	}
}
