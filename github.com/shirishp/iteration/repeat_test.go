package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {
	got := Repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)

	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 6)
	fmt.Println(repeat)
	//Output aaaaaa
}
