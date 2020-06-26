package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type ConfigurableSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep(){
	//.sleep references a value in the struct that is itself a function
	// when called, it gets the
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration){
	s.durationSlept = duration
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type CountdownOperationsSpy struct{
	Calls []string
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}


func main() {
	out := os.Stdout
	Sleeper := &ConfigurableSleeper{3 * time.Second, time.Sleep}
	Countdown(out, Sleeper)
}

func Countdown(out io.Writer, s Sleeper){
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)

	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}

func (s *CountdownOperationsSpy) Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte)(n int, err error){
	s.Calls = append(s.Calls, write)
	return
}
const write = "write"
const sleep = "sleep"