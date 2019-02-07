// This is a comment

/*
 This is multiline Comment
*/

package main

// import allows us to use code from other packages by using
import (
	"fmt" // The format package provides formatting for input and output (printing)
)

// Functions start with func and surround the code with { }
// main is the function that is executed when you execute your program
func main() {
	//The Println function of the fmt package is used to write text to the standard string output, which
	// is surrounded by double quotes and a newline to the screen
	fmt.Println("Hello World")

	// You can get a description of a function by typing `godoc fmt Println`
	// for example in the terminal
}

/*
Package:-
	- Every Go program is made up of packages. Programs start running in package main.
	- Every Go program starts with a package declaration which provides a way for use to reuse code
	- Package declaration at top of every source file
	- All Executable are in package `main`
	- Convention: package name == last name of import path (import path `math/rand` => package `rand`)
	- Upper case identifier: exported (visible from other packages)
	- Lower case identifier: private (not visible from other packages)
	- Usually, non standard lib packages are namespaced using a web url.Eg. import "github.com/mattetti/goRailsYourself/crypto".

Exported Name:-
	- After importing a package, you can refer to the names it exports (meaning variables, methods and functions that are available from outside of the package).
	- a name is exported if it begins with a CAPITAL LETTER. Else, it is not exported.
	- Any "unexported" names are not accessible from outside the package.
*/
