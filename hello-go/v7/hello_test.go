package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jamiul", "")
		want := "Hello, Jamiul"
		assertCorrectMessage(t, got, want)
	})

	// saying 'Hello, World' when an empty string is supplied
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	// say hello in Spanish
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Jamiul", "Spanish")
		want := "Hola, Jamiul"
		assertCorrectMessage(t, got, want)
	})

	// say hello in French
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Jamiul", "French")
		want := "Bonjour, Jamiul"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}