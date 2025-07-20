package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jamiul")
	want := "Hello, Jamiul"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}