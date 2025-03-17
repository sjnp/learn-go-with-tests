package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("with a name", func(t *testing.T) {
		got := Hello("Foo", "")
		want := "Hello, Foo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Foo", "Spanish")
		want := "Hola, Foo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Foo", "Japanese")
		want := "Konichiwa, Foo"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
