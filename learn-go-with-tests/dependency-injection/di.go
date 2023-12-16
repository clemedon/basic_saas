package main

import (
	"bytes"
	"fmt"
)

func Greet(dest *bytes.Buffer, name string) {
	fmt.Fprintf(dest, "Hello, %s", name)
}
