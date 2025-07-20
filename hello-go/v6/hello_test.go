package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("to a person", func(t *testing.T) {
		got := Hello("Jamiul", "")
		want := "Hello, Jamiul"
		assertCorrectMessage(t, got, want)
	})

	// Test for an empty name
	t.Run("to an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// Test for spanish greeting
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jamiul", "Spanish")
		want := "Hola, Jamiul"
		assertCorrectMessage(t, got, want)
	})

	// Test for french greeting
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jamiul", "French")
		want := "Bonjour, Jamiul"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}