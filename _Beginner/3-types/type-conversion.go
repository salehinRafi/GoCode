package main

import "fmt"

func TypeConversion() {

	// T(v) converts the value v to the type T
	var i int = 42
	fmt.Printf("%T\n", i)
	var f float64 = float64(i)
	fmt.Printf("%T\n", f)
	var u uint = uint(f)
	fmt.Printf("%T\n", u)
}

/*
	Go assignment between items of different type requires an explicit conversion which means that you manually need to convert types if you are passing a variable to a function expecting another type.
*/
