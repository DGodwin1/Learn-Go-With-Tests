package Arrays_and_Slices

import "testing"

func AssertCorrect(t *testing.T, got, want int, n []int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d want %d from %v", got, want, n)
	}

}

func TestSum(t *testing.T) {
	t.Run("test with five elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		AssertCorrect(t, got, want, numbers)
	})

	t.Run("test with any number of elements", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		AssertCorrect(t, got, want, numbers)
	})

}
