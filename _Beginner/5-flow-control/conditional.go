package main

import (
	"fmt"
	"math"
	"runtime"
)

// Conditional : Example in Golang
func Conditional() {

	// If Statement

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You Can Drive")
	} else {
		fmt.Println("You Can't Drive")
	}

	// If with a short statement
	var lim float64 = 10
	if v := math.Pow(2, 3); v < lim {
		fmt.Printf("%g\n", v)
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though

	// Else-If Statement

	if yourAge >= 16 {
		fmt.Println("You Can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You Can Have Fun")
	}

	// Switch statements

	switch yourAge {
	case 16:
		fmt.Println("Go Drive")
		// cases break automatically, no fallthrough by default
	case 18:
		fmt.Println("Go Vote")
	default:
		fmt.Println("Go Have Fun")
	}

	// as with for and if, you can have an ASSIGNMENT STETEMENT before the switch value
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Go Drive")
	}

	// you can also make comparisons in switch cases
	number := 42
	switch {
	case number < 42:
		fmt.Println("Smaller")
	case number == 42:
		fmt.Println("Equal")
	case number > 42:
		fmt.Println("Greater")
	}

	// cases can be presented in comma-separated lists
	var char byte = '?'
	switch char {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("Should escape")
	}
}
