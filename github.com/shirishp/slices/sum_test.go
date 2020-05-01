package slices

import "testing"

func TestSum(t *testing.T) {
	numbers := [] int {1,2,3,4}
	got := Sum(numbers)
	expected := 10

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
