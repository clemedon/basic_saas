package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(dest io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(dest, i)
	}
	fmt.Fprint(dest, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
