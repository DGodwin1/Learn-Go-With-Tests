package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type DefaultSleeper struct {
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func (d *DefaultSleeper) Sleep(){
	time.Sleep(1 * time.Second)
}



func main(){
	out := os.Stdout
	sleeper := &DefaultSleeper{}
	Countdown(out, sleeper)
}

func Countdown(out io.Writer, s Sleeper){
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}
