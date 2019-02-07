/*
	Two MAIN reasons to use a pointer receiver:-
	1 - to avoid copying the value on each method call (more efficient if the value type is a large struct).
	2 - the method can modify the value that its receiver points to.
*/

package main

import (
	"fmt"
	"math"
)

// 1st reason example:-
// Remember that Go passes everything by value, meaning that when Greeting() is defined on the value type, every time you call Greeting(), you are copying the User struct. Instead when using a pointer, only the pointer is copied (cheap).

type User struct {
	FirstName, LastName string
}

func (u *User) Greeting() string { // if we dont use pointer, we are copying the User struct.
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

// 2nd reason example:-

type VertexPointer struct {
	X, Y float64
}

// Scale : Scale() has to be defined on a pointer since it does modify the receiver.
// Scale() resets the values of the X and Y fields.
// here we have function with [(receiver) function-name (argument/parameter)]
func (v *VertexPointer) Scale(f float64) { //
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs : Abs() could be defined on the value type or the pointer since the method doesn’t modify the receiver value (the vertex).
func (v *VertexPointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func PointerReceiver() {

	// 1st reason example
	u := &User{"Salehin", "Rafi"}

	fmt.Println("1st-reason: ", u.Greeting()) // Call the method Greeting for User

	// 2nd reason example
	v := &VertexPointer{3, 4}
	v.Scale(5)                            // Call the method Scale for VertexPointer to modify value
	fmt.Println("2nd-reason", v, v.Abs()) // Call the method Abs() for VertexPointer
}
