package array_and_slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice and array", func(t *testing.T) {
		//numbers := [5]int{1, 2, 3, 4, 5}
		//same as
		//numbers := [...]int{1, 2, 3, 4, 5}

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}
	//expected := "deepEq not safe type"

	//careful when using it (deepEqual)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run("Make the sums of slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})
}
