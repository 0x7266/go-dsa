package isprime

import "testing"

func TestOne(t *testing.T) {
	expected := false
	got := IsPrime(1)

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func TestTwo(t *testing.T) {
	expected := true
	got := IsPrime(2)

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func TestSeven(t *testing.T) {
	expected := true
	got := IsPrime(7)

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func TestNine(t *testing.T) {
	expected := false
	got := IsPrime(9)

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func TestSixtyFour(t *testing.T) {
	expected := false
	got := IsPrime(64)

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func TestSixtySeven(t *testing.T) {
	expected := true
	got := IsPrime(67)

	if expected != got {
		t.Errorf("expected %t but got %t", expected, got)
	}
}
