package main

import "fmt"

/*
i. Sender
	- A sender can close a channel to indicate that no more values will be sent.
ii. Receiver
	- Receivers can test whether a channel has been closed by assigning a second parameter to the receiver expression.
iii. Expression
	- v, ok := <-ch , where v as a sender, ok as a receiver. (ok is false if there are no more values to receive and the channel is closed.)

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren’t like files; you don’t usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // Closing a channel indicates that no more values will be sent on it.
}

// RangeOverChannel :
// - Example of how to iterate over value received from channel.RangeOverChannel
// - Combination with example of closing channel
func RangeOverChannel() {
	// 1st Example
	// like maps and slices, channels must be created before use:
	c := make(chan int, 10)
	go fibonacci(cap(c), c) //cap built-in function returns the capacity of v, according to its type. In this case return as channel type
	for i := range c {      //receives values from the channel repeatedly until it is closed.
		fmt.Println(i)
	}

	// 2nd Example
	queue := make(chan string, 2)
	//iterate over 2 values in the queue channel.
	queue <- "one"
	queue <- "two"
	close(queue) // Closing a channel indicates that no more values will be sent on it.

	// range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
