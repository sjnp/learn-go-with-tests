package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 3)
	expected := "aaa"

	assertCorrectMessage(t, expected, got)
}

func assertCorrectMessage(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {

	b.Run("slow repeat", func(b *testing.B) {
		// This will run 10^7 times to find average ns per op
		for b.Loop() {
			Repeat("a", 5)
		}
	})

	b.Run("fast repeat", func(b *testing.B) {
		// This will run 10^7 times to find average ns per op
		for b.Loop() {
			FastRepeat("a", 5)
		}
	})
}

func ExampleRepeat() {
	result := Repeat("a", 3)
	fmt.Println(result)
	// Output: aaa
}
