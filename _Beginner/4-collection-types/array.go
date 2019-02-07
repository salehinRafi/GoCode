package main

import "fmt"

// Array : Example in Go
// Arrays have a fixed size of the same type.
// [n]T is an array of n length of type T.
func Array() {
	var a [5]int
	b := [2]int{1, 2}         // shorthand declaration
	fmt.Println("emp:", a, b) // default an array is zero-valued

	a[3] = 42 // set elements
	fmt.Println("set:", a)
	i := a[3] // read elements by supplying index number
	fmt.Println("get:", i)
	fmt.Println("len:", len(a)) // length of an array

	c := [5]int{1, 2, 3, 4, 5} // declare and initialize
	fmt.Println("dcl:", c)

	// build multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	d := [...]int{1, 2} // elipsis -> Compiler figures out array length
	fmt.Println("elipsis: ", d)
}
