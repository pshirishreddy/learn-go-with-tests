package slices

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	a := [] int {1,2}
	b := [] int {2,3}
	got := SumAll(a, b)
	expected := [] int {3, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v got %v", expected, got)
	}
}
