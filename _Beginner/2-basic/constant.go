package main

import "fmt"

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

// Constant : Constant declaration in Go
func Constant() {
	/*
	 Constants can only be character, string, boolean, or numeric values and cannot be declared using the := syntax. An untyped constant takes the type needed by its context.
	*/
	const Greeting = "Selamat Datang"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
}
