package main

import (
	"fmt"
)

// Variable : Variable declaration in Go
func Variable() {

	// Variables are statically typed, which means their type can't change (immutable).
	// Variable names must start with a letter and may contain letters, numbers or the _
	// `var` statement can be at package or function level

	// declaration without initialization
	var a int
	a = 1
	fmt.Println(a)

	// declaration with initialization
	// type omitted, will be inferred
	var e = 42
	fmt.Println(e)

	// shorthand, only in func bodies, omit var keyword, type is always implicit
	foo := 42
	fmt.Println(foo)

	// Multiple declaration at once
	var c, d int = 42, 1302
	fmt.Println(c, d)

	var (
		varA = 2
		varB = "Hello"
	)
	fmt.Println(varA, varB)

	// Variable can contain any type, including function
	action := func() {
		fmt.Println("Variable as function")
	}
	action()
}
