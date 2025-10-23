package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to person", func(t *testing.T) {
		name := "Jenny"
		got := Hello(name, english)
		want := fmt.Sprintf("Hello, %s!", name)

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		emptyName := ""
		got := Hello(emptyName, english)
		want := fmt.Sprintf("Hello, World!")

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to person in Spanish", func(t *testing.T) {
		name := "Jenny"
		lang := "Spanish"
		got := Hello(name, lang)
		want := fmt.Sprintf("Hola, %s!", name)

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
