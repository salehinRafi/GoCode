package main

import (
	"fmt"
	"time"
)

/*
GoSelect :-

	- `select` statement lets a goroutine wait on multiple communication operations.
	- A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/
func GoSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		// use `select` to await both of these values simultaneously, printing each one as it arrives.
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
