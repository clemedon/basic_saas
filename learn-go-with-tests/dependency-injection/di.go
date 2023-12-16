package main

import (
	"fmt"
	"io"
	"os"
)

// 'fmt.Fprintf' allows you to pass in an 'io.Writer' which both 'os.Stdout' and
// 'bytes.Buffer' implement, by using a more general purpose interface we can
// use it in both tests and our application.

func Greet(dest io.Writer, name string) {
	fmt.Fprintf(dest, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
