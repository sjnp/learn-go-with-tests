package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 4)
	expected := "aaaa"

	if result != expected {
		t.Errorf("got %q but expected %q", result, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat() {
	result := Repeat("f", 2)
	fmt.Println(result)
	// Output: ff
}

func BenchmarkStandardRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardRepeat("a", b.N)
	}
}
