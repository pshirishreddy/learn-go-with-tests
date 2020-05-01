package slices

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, expected []int, got []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
	t.Run("make sum of slices", func(t *testing.T) {
		a := [] int {1,2,3}
		b := [] int {3,4,5}
		got := SumAllTails(a, b)
		expected := [] int {5, 9}
		checkSums(t,expected, got)
	})

	t.Run("make sum with empty slices", func(t *testing.T) {
		a := [] int {}
		b := [] int {3,4,5}
		got := SumAllTails(a, b)
		expected := [] int {0, 9}
		checkSums(t, expected, got)
	})
}


