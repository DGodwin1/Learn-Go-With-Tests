package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	// buffer means we can deal with std.out dependencies related
	// to printing back results to a user. Which is handy.
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := SpySleeper{}

		Countdown(buffer, &spySleeper)

		got := buffer.String()

		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		// Faster than actual sleeping because you're not counting seconds.
		// You're just incrementing a count. And because you call Sleep()
		// four times in the process of executing countdown, then spySLeeper.Calls
		// should equal 4 to reflect that sleeping 'sort of' took place.
		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		// Recall that Countdown takes an IOwriter and a Sleeper.
		// The SpySleepPrinter type implements both Writer and Sleeper.
		// When write is called on spySleepPrinter, we're actually not writing
		// anything, we're running the append on the slice.
		// Same is true for Sleep(). We don't actually do any sleeping. We just append
		// the string to slice.
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v but got %v", want, spySleepPrinter.Calls)
		}
	})
}
