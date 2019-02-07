package main

import "fmt"

//Operator : Operator in Golang
func Operator() {

	/*
		### Arithmetic

		| Operator | Description         |
		| -------- | ------------------- |
		| `+`      | addition            |
		| `-`      | subtraction         |
		| `*`      | multiplication      |
		| `/`      | quotient            |
		| `%`      | remainder           |
		| `&`      | bitwise and         |
		| `\|`     | bitwise or          |
		| `^`      | bitwise xor         |
		| `&^`     | bit clear (and not) |
		| `<<`     | left shift          |
		| `>>`     | right shift         |

		### Comparison

		| Operator | Description           |
		| -------- | --------------------- |
		| `==`     | equal                 |
		| `!=`     | not equal             |
		| `<`      | less than             |
		| `<=`     | less than or equal    |
		| `>`      | greater than          |
		| `>=`     | greater than or equal |

		### Logical

		| Operator | Description |
		| -------- | ----------- |
		| `&&`     | logical and |
		| `\|\|`   | logical or  |
		| `!`      | logical not |

		### Other

		| Operator | Description                                    |
		| -------- | ---------------------------------------------- |
		| `&`      | address of / create pointer                    |
		| `*`      | dereference pointer                            |
		| `<-`     | send / receive operator (see 'Channels' below) |
	*/
	// Logical Operators

	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

}
