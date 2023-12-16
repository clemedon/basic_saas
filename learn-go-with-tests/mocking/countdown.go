package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(dest io.Writer) {
	fmt.Fprint(dest, "3")
}

func main() {
	Countdown(os.Stdout)
}
