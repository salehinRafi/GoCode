package main

import (
	"fmt"
	"time"
)

// Range : Example in Go
// Operations on Arrays and Slices
func Range() {
	// `len(a)` gives you the length of an array/a slice. It's a built-in function, not a attribute/method on the array.

	// loop over an array/a slice/a map

	var a = []int{1, 2, 3}
	for i, e := range a {
		// i is the index, e the element
		fmt.Println("index of:", i, "with element:", e)
	}

	// if you only need e:
	for _, e := range a {
		// e is the element
		fmt.Println("Element only:", e)
	}

	// ...and if you only need the index
	for i := range a {
		fmt.Println("Index only:", i)
	}

	// In Go pre-1.4, you'll get a compiler error if you're not using i and e.
	// Go 1.4 introduced a variable-free form, so that you can do this
	for e := range time.Tick(time.Second) {
		// do it once a sec
		fmt.Println(e)
		break // try remove this and run
	}

}
