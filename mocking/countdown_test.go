package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T){
	// buffer means we can deal with std.out dependencies related
	// to printing back results to a user. Which is handy.
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()

	want := `3
2
1
Go!`

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}