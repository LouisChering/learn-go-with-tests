package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Louis", "")
		want := "Hello, Louis"
		if got != want {
			assertMessage(t, got, want)
		}
	})
	t.Run("say hello to world when empty string provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jodie", "French")
		want := "Bonjour, Jodie"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
