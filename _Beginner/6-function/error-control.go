package main

import "fmt"

// ErrorControl - Example of defer,panic, recover in go
func ErrorControl() {

	// Defer executes a function after the inclosing (in this case is Defer function) function finishes
	// Defer can be used to keep functions together in a logical way
	// but at the same time execute one last as a clean up operation
	// Ex. Defer closing a file after we open it and perform operations

	defer printTwo() // execute at the end of this function (last)
	printOne()

	// Use recover() to catch a division by 0 error

	fmt.Println("Catch error:", safeDiv(3, 0))
	fmt.Println("Success:", safeDiv(3, 2))

	// We can catch our own errors and recover with panic & recover

	demPanic()

}

// Used to demonstrate defer
func printOne() { fmt.Println("PrintOne") }

func printTwo() { fmt.Println("PrintTwo") }

// If an error occurs we can catch the error with recover and allow
// code to continue to execute

func safeDiv(num1, num2 int) int {
	// this defer will execute at the of function
	defer func() {
		fmt.Println("safeDiv recover:", recover()) // catch the error with recover()
	}() // use () as expression in defer for function call
	solution := num1 / num2
	return solution
}

// Demonstrate how to call panic and handle it with recover
func demPanic() {

	defer func() {

		// If I didn't print the message nothing would show

		fmt.Println("Recover:", recover())

	}()
	panic("PANIC")

}
