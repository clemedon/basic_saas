package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type RealSleeper struct{}

func (d *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(dest io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(dest, i)
		sleeper.Sleep()
	}
	fmt.Fprint(dest, finalWord)
}

func main() {
	sleeper := &RealSleeper{}
	Countdown(os.Stdout, sleeper)
}
