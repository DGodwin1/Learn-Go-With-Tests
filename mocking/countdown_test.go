package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T){
	// buffer means we can deal with std.out dependencies related
	// to printing back results to a user. Which is handy.
	buffer := &bytes.Buffer{}
	spySleeper := SpySleeper{}

	Countdown(buffer, &spySleeper)

	got := buffer.String()

	want := `3
2
1
Go!`

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}

	// Faster than actual sleeping because you're not counting seconds.
	// You're just iterating up a count. And because you call Sleep()
	// four times in the process of executing countdown, then spySLeeper.Calls
	// should equal 4 to reflect that sleeping 'sort of' took place.
	if spySleeper.Calls != 4{
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
