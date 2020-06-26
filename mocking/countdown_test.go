package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

func TestConfigurableSleeper(t *testing.T){
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime{
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
