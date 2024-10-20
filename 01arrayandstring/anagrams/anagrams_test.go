package anagrams

import "testing"

func Test12y(t *testing.T) {
	expected := true
	got := Anagrams("12y", "y12")

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func Test1y2x6c5u(t *testing.T) {
	expected := true
	got := Anagrams("1y2x6c5u", "x6c5u1y2")

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func TestAaabaacc(t *testing.T) {
	expected := false
	got := Anagrams("Aaabaacc", "Aaacbbaa")

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}
