package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(dest io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(dest, i)
	}
	fmt.Fprint(dest, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
