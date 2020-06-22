package Arrays_and_Slices

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T){
	t.Run("Sum up varied size of args", func(t *testing.T){
		got := SumAll([]int{1,2,3}, []int{1,2})
		want := []int{6, 3}

		if !reflect.DeepEqual(got, want){ //not type safe but let's you check slices.
			t.Errorf("got %d wanted %d", got, want)
		}
	})

	t.Run("Test the addition of tail elements", func(t *testing.T){
		got := SumAllTails([]int{1,2,3}, []int{1,2})
		want := []int{5,2}
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %d wanted %d", got, want)
		}
	})

	t.Run("Test adding a zero tail", func(t *testing.T){
		got := SumAllTails([]int{1}, []int{1})
		// Here we'll say that it should give back 0,0, but it could just easily abort and say 'nope'
		// because it can't operate.
		want := []int{0,0}
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %d wanted %d", got, want)
		}
	})
	t.Run("Test adding a zero tail and a normal tail", func(t *testing.T){
		got := SumAllTails([]int{1}, []int{1, 9})
		// Here we'll say that it should give back 0,0, but it could just easily abort and say 'nope'
		// because it can't operate.
		want := []int{0,9}
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %d wanted %d", got, want)
		}
	})
}
