package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	result := Add(1, 1)
	expected := 2

	assertCorrectMessage(t, expected, result)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, expected, got int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}
