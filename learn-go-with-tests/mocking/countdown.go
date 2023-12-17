package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const countdownStepsDelay = 1
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(dest io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(dest, i)
		sleeper.Sleep()
	}
	fmt.Fprint(dest, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{countdownStepsDelay * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
