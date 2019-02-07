package main

import (
	"fmt"
)

// Deferred : Go has a special statement called `defer` that schedules a function call to be run after the function completes.
// Deferred functions are executed in LIFO order.
func Deferred() {
	defer second()
	first()

	// print first() function 1st. Then second() will be executed then after all the code inside the codeblock is complete.
	defer fmt.Println("")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}

/*
defer statement is often used with paired operations like open and close, connect and disconnect, or lock and unlock to ensure that resources are released in all cases, no matter how complex the control flow

advantage: -

- It keeps our Close call near our Open call so it's easier to understand.
- If our function had multiple return statements (perhaps one in an if and one in an else), Close will happen before both of them.
- Deferred Functions are run even if a runtime panic occurs.
- Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
- You can put multiple functions on the "deferred list".
*/
