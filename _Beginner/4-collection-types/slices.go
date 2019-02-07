package main

import (
	"fmt"
)

// Slices : Example slice in go
// Slices have a dynamic size, unlike arrays.
func Slices() {

	var a []int               // declare a slice - similar to an array, but length is unspecified
	fmt.Println(a)            // produce zero value or nil (A nil slice has a length and capacity of 0.)
	var b = []int{1, 2, 3, 4} // declare and initialize a slice (backed by the array given implicitly)
	c := []int{1, 2, 3, 4}    // shorthand declaration
	fmt.Println(c)
	chars := []string{0: "a", 2: "c", 1: "b"} // ["a", "b", "c"]
	fmt.Println(chars)

	/* Slicing a slice:-
	-  d = a[lo:hi] , creates a slice (view of the array) from index lo to hi-1
	*/
	var d = b[1:4] // slice from index 1 to 3
	fmt.Println("slice from index 1 to 3:", d)
	var e = b[:3] // missing low index implies 0 index
	fmt.Println("missing low index implies 0 index:", e)
	var f = b[3:] // missing high index implies len(a)
	fmt.Println("missing high index implies len(a):", f)
	a = append(a, 17, 3) // append items to slice a
	fmt.Println(chars)
	g := append(a, b...) // concatenate two slices a and b using ellipsis `...`
	fmt.Println("concatenate slices a and b", g)

	// create a slice with `make` before use
	// make function allocates a zeroed array and returns a slice that refers to that array.
	var h = make([]byte, 5, 5) // first arg length, second capacity
	fmt.Println("first arg length, second capacity:", h)
	var i = make([]byte, 5) // capacity is optional
	fmt.Println("capacity is optional:", i)

	// create a slice from an array
	x := [3]string{"Salehin", "Rafi", "CodeLads"}
	s := x[:] // a slice referencing the storage of x
	fmt.Println("create a slice from an array:", s)
}
