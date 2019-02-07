package main

import "fmt"

func VariadicFunction() {
	// the function which accepts a variable number of arguments.

	fmt.Println(adder(1, 2, 3)) // 6
	fmt.Println(adder(9, 9))    // 18

	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...)) // 60
}

// By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
// The function is invoked like any other function except we can pass as many arguments as we want.
func adder(args ...int) int {
	total := 0
	//It avoids having to declare all the variables for the returns values.
	//It is called the blank identifier.
	// Eg: _, y, _ := coord(p)  // coord() returns three values; only interested in y coordinate
	for _, v := range args { // Iterates over the arguments whatever the number.
		total += v
	}
	return total
}
