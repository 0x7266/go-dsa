package pairsum

import (
	"reflect"
	"testing"
)

func TestEight(t *testing.T) {
	got := PairSum([]int{3, 2, 5, 4, 1}, 8)
	expected := []int{0, 2}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestEighteen(t *testing.T) {
	got := PairSum([]int{9, 9}, 18)
	expected := []int{0, 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestTwelve(t *testing.T) {
	got := PairSum([]int{6, 4, 2, 8}, 12)
	expected := []int{1, 3}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
