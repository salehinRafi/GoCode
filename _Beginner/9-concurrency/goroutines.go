package main

import (
	"fmt"
	"time"
)

// just a function (which can be later started as a goroutine)
func say(s string) {
	for i := 0; i < 5; i++ {

		time.Sleep(100 * time.Millisecond) // Pause the function for 1 second to allow other functions to execute
		fmt.Println(s)
	}
}

// Goroutines = goroutine is a lightweight thread managed by the Go runtime.
// Goroutines run in the same address space,so access to shared memory must be synchronized.
// the order isn't guaranteed!)
func Goroutines() {

	// A go routine is a function that runs in parallel with other functions
	// We define one by using go followed by the function name
	go say("world")
	say("hello")

	// Wait for the timer to make sure the go routine has time to
	// finish otherwise the program would end before that happens
	time.Sleep(time.Millisecond * 11000)

	// using an anonymous inner function in a goroutine
	go func(x int) {
		// function body goes here
	}(42)

}
