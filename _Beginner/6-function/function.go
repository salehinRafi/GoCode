package main

import (
	"fmt"
)

// Function : Example of Simple Function
func Function() {

	// a simple function
	functionName1()

	// function with parameters (again, types go after identifiers)
	functionName2("Salehin Rafi", 26)

	// multiple parameters of the same type
	functionName3(1, 2)

	// return type declaration
	var i = functionName4()
	fmt.Println(i)

	// return multiple values at once
	var a, b = returnMulti()
	fmt.Println(a, b)

	// return multiple named results simply by return
	var c, d = returnMulti2()
	fmt.Println(c, d)

	// assign a function to a name as a value
	add := func(a, b int) int {
		return a + b
	}
	// use the name to call the function
	fmt.Println(add(3, 4))

}

// functionName1 : a simple function
func functionName1() {
	fmt.Println("Simple Function")
}

// functionName2 : function with parameters (again, types go after identifiers)
func functionName2(param1 string, param2 int) {
	fmt.Println(param1, param2)
}

// functionName3 : multiple parameters of the same type
func functionName3(param1, param2 int) {
	fmt.Print(param1, param2)
}

// functionName4 : return type declaration
func functionName4() int {
	return 42
}

// returnMulti : Can return multiple values at once
func returnMulti() (int, string) {
	return 42, "foobar"
}

// returnMulti2 : return multiple named results simply by return
func returnMulti2() (n int, s string) {
	n = 42
	s = "foobar"
	// n and s will be returned
	return
}
