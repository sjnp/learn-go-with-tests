package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John")

	got := buffer.String()
	expected := "Hello, John"

	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}
