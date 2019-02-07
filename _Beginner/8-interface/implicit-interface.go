package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

// Implicit : Implicit interfaces decouple implementation packages from the packages that define the interfaces, neither depends on the other.
func Implicit() {
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
