package mostfrequentchar

import "testing"

func TestDavid(t *testing.T) {
	expected := 'd'
	got := MostFrequentChar("david")

	if expected != got {
		t.Errorf("expected %c but got %c", expected, got)
	}
}

func TestMississippi(t *testing.T) {
	expected := 'i'
	got := MostFrequentChar("mississippi")

	if expected != got {
		t.Errorf("expected %c but got %c", expected, got)
	}
}
