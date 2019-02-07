package main

import "fmt"

// Printing : Example of fmt package that use for printing.
func Printing() {

	const pi float64 = 3.14159265359
	var isOver40 bool = false

	// Printf is used for format printing (%f is for floats)
	fmt.Printf("%f \n", pi)

	// You can also define the decimal precision of a float
	fmt.Printf("%.3f \n", pi)

	// %T prints the data type
	fmt.Printf("%T \n", pi)

	// %t prints booleans
	fmt.Printf("%t \n", isOver40)

	// %d is used for integers
	fmt.Printf("%d \n", 100)

	// %b prints in binary
	fmt.Printf("%b \n", 100)

	// %c prints the character associated with a keycode
	fmt.Printf("%c \n", 44)

	// %x prints in hexcode
	fmt.Printf("%x \n", 17)

	// %e prints in scientific notation
	fmt.Printf("%e \n", pi)

	fmt.Println("Hello, Salehin, Ziema") // basic print, plus newline
	p := struct{ X, Y int }{17, 2}
	fmt.Println("My point:", p, "x coord=", p.X)       // print structs, ints, etc
	s := fmt.Sprintln("My point:", p, "x coord=", p.X) // print to string variable

	fmt.Println(s)

	fmt.Printf("%d hex:%x bin:%b fp:%f sci:%e", 17, 17, 17, 17.0, 17.0) // c-ish format
	s2 := fmt.Sprintf("%d %f", 17, 17.0)                                // formatted print to string variable

	fmt.Println(s2)

	hellomsg := `
 	"Hello" in Chinese is 你好 ('Ni Hao')
 	"Hello" in Hindi is नमस्ते ('Namaste')
	` // multi-line string literal, using back-tick at beginning and end

	fmt.Println(hellomsg)
}
