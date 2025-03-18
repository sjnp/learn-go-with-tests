package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("got %d but expected %d from %v", result, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3, 4}
	numbers2 := []int{5, 6, 7}
	result := SumAll(numbers1, numbers2)
	expected := []int{10, 18}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %d but expected %d from %v and %v", result, expected, numbers1, numbers2)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %v but expected %v", result, expected)
		}
	}

	t.Run("with all slice", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4}
		numbers2 := []int{5, 6, 7}
		numbers3 := []int{8, 9}
		result := SumAllTails(numbers1, numbers2, numbers3)
		expected := []int{9, 13, 9}
		checkSums(t, result, expected)
	})

	t.Run("with some empty slice", func(t *testing.T) {
		var numbers1 []int
		numbers2 := []int{5, 6, 7}
		result := SumAllTails(numbers1, numbers2)
		expected := []int{0, 13}
		checkSums(t, result, expected)
	})
}
