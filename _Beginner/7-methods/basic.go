/*
Go does not have classes. However, you can define methods on types.

A method is a function with a receiver argument.
func (receiver) method-name(arg) return type {
	. . .
}, where receiver is identifiers and types [again, types go after identifiers].
, arg is function argument/parameter with identifiers and types [again, types go after identifiers]

The receiver appears in its own argument list between the func keyword and the method name.
*/
package main

import (
	"fmt"
	"math"
)

/** 1. Method on Struct **/

type Vertex struct {
	X, Y float64
}

// Abs : method with a receiver of type Vertex (struct) named v with a return value with type float64
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/** 2. Method on Non-Struct **/

type MyFloat float64

// Abs2 : method with a receiver of type MyFloat (non-struct) named f with a return value with type float64
func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Basic() {

	/** 1. Method on Struct **/
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	/** 2. Method on Non-Struct
		- methods were defined on the value types (MyFloat).
	 **/
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2())
}
