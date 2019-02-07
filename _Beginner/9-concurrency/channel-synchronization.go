package main

import (
	"fmt"
	"time"
)

// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we’re done.
	done <- true
}

// ChannelSynchronization : Example to use channels to synchronize execution across goroutines.
func ChannelSynchronization() {
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
	// If removed the `<-done` line from this program, the program would exit before the worker even started.
}
