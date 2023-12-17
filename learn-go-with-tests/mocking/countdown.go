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

const finalWord = "Go!"
const countdownStart = 3

func Countdown(dest io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(dest, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(dest, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
