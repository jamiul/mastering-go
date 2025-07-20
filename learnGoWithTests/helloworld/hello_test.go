package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("to a person", func(t *testing.T) {
		got := Hello("Jamiul", "")
		want := "Hello, Jamiul"
		assertCorrectMessage(t, got, want)
	})

	// for the empty string case
	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "World")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// for the spanish case
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jamiul", "Spanish")
		want := "Hola, Jamiul"
		assertCorrectMessage(t, got, want)
	})
	// for the french case
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jamiul", "French")
		want := "Bonjour, Jamiul"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
