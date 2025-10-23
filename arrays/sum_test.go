package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum of any numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		expected := 6

		if got != expected {
			t.Errorf("got %d but expected %d given %v", got, expected, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5})
	expected := []int{6, 9}

	checkSums(t, got, expected)
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{4, 5}, []int{})
	expected := []int{5, 5, 0}

	checkSums(t, got, expected)
}

func checkSums(tb testing.TB, got []int, expected []int) {
	// mark that this func is a helper function
	// the line that the test fail will be at the line that call this function
	tb.Helper()

	if !slices.Equal(got, expected) {
		tb.Errorf("got %v but expected %v", got, expected)
	}
}
