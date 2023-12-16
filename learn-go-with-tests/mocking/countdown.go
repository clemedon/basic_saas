package main

import (
	"bytes"
	"fmt"
)

func Countdown(dest *bytes.Buffer) {
	fmt.Fprint(dest, "3")
}
