package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func main(){
	out := os.Stdout
	Countdown(out)
}

func Countdown(out io.Writer){
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1*time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1*time.Second)
	fmt.Fprint(out, finalWord)
}
