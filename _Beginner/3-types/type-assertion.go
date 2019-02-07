package main

import (
	"fmt"
	"time"
)

func TypeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // produce panic error , because assertion always return 2 value: underlying and boolean
	// /fmt.Println(f)

	//In the example below, the timeMap function takes a value and if it can be asserted as a map of interfaces{} keyed by strings, then it injects a new entry called “updated_at” with the current time value.

	foo := map[string]interface{}{
		"Matt": 42,
	}
	timeMap(foo)
	fmt.Println(foo)
}

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

/*
	If you have a value and want to convert it to another or a specific type (in case of interface{}).
	A type assertion takes a value and tries to create another version in the specified explicit type.

	A type assertion provides access to an interface value's underlying concrete value.
	t := i.(T)
	This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

	If i does not hold a T, the statement will trigger a panic.

	To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

	t, ok := i.(T)
	If i holds a T, then t will be the underlying value and ok will be true.

	If not, ok will be false and t will be the zero value of type T, and no panic occurs.

	Note the similarity between this syntax and that of reading from a map.
*/
