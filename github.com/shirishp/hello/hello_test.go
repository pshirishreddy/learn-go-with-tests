package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Shirish")
	want := "Hello, Shirish"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}