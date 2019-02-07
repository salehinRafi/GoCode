package main

import "fmt"

// Loops : Example in Golang
func Loops() {
	// There's only `for`, no `while`, no `until`
	for i := 1; i < 2; i++ {
		fmt.Println("Loop")
	}
	// Can act as a while loop. you can omit semicolons if there is only a condition or as a for loop without pre/post statements
	j := 1
	for j < 2 {
		fmt.Println("while-loop")
		break
	}

	// infinity loop - you can omit the condition ~ while (true)
	for {
		fmt.Println("Without condition ")
		break
	}

	// Loop Control Statements
	/*
		i. break 	 : = When a break statement is encountered inside a loop, the loop is immediately
		terminated and the program control resumes at the next statement following the loop.

		ii. continue := a continue statement forces the next iteration of the loop to take place, skipping
		any code in between.

		ii.goto 	 := goto statement in Go programming language provides an unconditional jump from the goto to a labeled statement in the same function.

		syntax:
			goto label;
			..
			.
			label: statement;
	*/

	// continue control
	var a = 8

	for a < 10 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			continue
		}
		fmt.Printf("value of a (continue): %d\n", a)
		a++
	}

	// goto
	var c = 5

LOOP:
	for c < 10 {
		if c == 6 {
			/* skip the iteration */
			c = c + 1
			fmt.Println("Execute goto control")
			goto LOOP
		}
		fmt.Printf("value of c (goto): %d\n", c)
		c++
	}
}
