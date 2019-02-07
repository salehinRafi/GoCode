package main

import "fmt"

// DataType : Data Types in GO
func DataType() {

	// Boolean types
	/* Data Type: bool
	 * Booleans can be either true or false
	 */
	var a bool = true
	var b bool = false
	fmt.Println("a:", a, "b:", b)

	/* Data Type: int
	 * An int is a positive or negative number without decimals
	 * Versions
	 * uint8 : unsigned  8-bit integers (0 to 255)
	 * uint16 : unsigned 16-bit integers (0 to 65535)
	 * uint32 : unsigned 32-bit integers (0 to 4294967295)
	 * uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	 * int8 : signed  8-bit integers (-128 to 127)
	 * int16 : signed 16-bit integers (-32768 to 32767)
	 * int32 : signed 32-bit integers (-2147483648 to 2147483647)
	 * int64 : signed 64-bit integers (-9223372036854775808 to
	 * 9223372036854775807)
	 */
	var age int64 = 26
	fmt.Println(age)

	/* Data Type: float
	 * A float is a number with decimals
	 * Versions
	 * float32 : 32 bit floating point numbers
	 * float64 : 64 bit floating point numbers
	 */

	var favNum float64 = 1.61803398875
	fmt.Println(favNum)

	/* Data Type: complex
	 * complex is used to construct a complex number with real and imaginary parts.
	 * Versions
	 * complex64: complex numbers which have float32 real and imaginary parts
	 * complex128: complex numbers with float64 real and imaginary parts
	 */
	c1 := complex(5, 7)
	fmt.Println(c1)
	c2 := 8 + 27i
	fmt.Println(c2)

	/* Data Type: string
	 * Strings are a series of characters between " or `
	 */

	first := "Salehin"
	last := "Rafi"
	name := first + " " + last
	fmt.Println("My name is", name)

	// You don't need to define the data type, nor do you need a semicolon
	// but you can use them

	randNum := 1
	fmt.Println(randNum)

	// You can't however assign a non-compatible type later

	// randNum = "Hello"

}
