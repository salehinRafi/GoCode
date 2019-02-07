package main

import "fmt"

// Closure : function as a closure
func Closure() {

	/*
		You can create a function in a function.
		It has access to the local variables of the containing function.
		A function like this with no local variables is a closure.

	*/
	// Create fibonacci
	gen := makeFibGen()
	for i := 0; i < 10; i++ {
		fmt.Println("fibonacci:", gen())
	}

	// Example 2
	getScope := scope()
	fmt.Println("Example 2:", getScope())

	// Example 3
	getInner, getOuter := outer()
	fmt.Println("Example 3 Inner:", getInner, "Example 3 Outer:", getOuter)

}

// Closures, lexically scoped: Functions can access values that were in scope when defining the function
// Closures: don't mutate outer vars, instead redefine them!
func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

// Another example
func scope() func() int {
	outerVar := 2
	foo := func() int {
		return outerVar
	}
	return foo
}

// Closures: don't mutate outer vars, instead redefine them!
func outer() (func() int, int) {
	outerVar := 2
	inner := func() int {
		outerVar += 99  // attempt to mutate outer_var from outer scope
		return outerVar // => 101 (but outer_var is a newly redefined
		// variable visible only inside inner)
	}
	return inner, outerVar // => 101, 2 (outer_var is still 2, not mutated by inner!)
}

/*
			5 Useful Ways to Use Closures in Go:=
			i. Isolating Data
					we want to create a function that has access to data that persists even after the function exits.
			ii. Wrapping functions and creating middleware
					We can not only create anonymous functions dynamically, but we can also pass functions as parameters to a function. For example, when creating a web server it is common to provide a function that processes a web request to a specific route.
			iii. Accessing data that typically isn’t available
					can also be used to wrap data inside of a function that otherwise wouldn’t typically have access to that data. For example, if you wanted to provide a handler access to a database without using a global variable.
			iv. Binary searching with the sort package
					Closure are also often needed to use packages in the standard library, such as the sort package. This package provides us with tons of helpful functions and code for sorting and searching sorted lists.
			v. Deferring work - similar to callback function in javascript

					doWork(a, b, function(result) {
	  					// use the result here
					});
					console.log("hi!");

					What we are essentially doing is telling our program to run the function doWork() with the variables a and b, and then our last argument is the function that we want it to run after that function completes. So when doWork() finishes, it then calls our function with there result of doWork().
*/
