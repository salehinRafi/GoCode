package main

import "fmt"

// fibonaci : Using recursion function to solve fibonaci
func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

// Example of recursion : Function calls itself
// factorial(3)
// 3 * factorial(2) == 3 * 2
// 2 * factorial(1) == 2 * 1
// factorial(0) == 1

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Recursion : Recursion is the process of repeating items in a self-similar way.
// To call a function inside the same function, then it is called a recursive function call.
// Be careful to define an exit condition from the function, otherwise it will go on to become an infinite loop.
func Recursion() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d", fibonaci(i))
	}

	// Calling a recursive function
	fmt.Println("recursive: ", factorial(3))
}
