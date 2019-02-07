package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to channel c
}

// Channel :
// - Channels are the pipes that connect concurrent goroutines.
// - Channels allow us to pass data (send & receive) between go routines with the channel operator, <-
func Channel() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // declaration and initialization (Channels are typed by the values they convey.)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c and assign to x, y
	fmt.Println(x, y, x+y)
}

// ChannelBuffering :
// - Buffered channels accept a limited number of values without a corresponding receiver for those values.
func ChannelBuffering() {

	messages := make(chan string, 2) // declaration and initialization with buffer length

	// we can send these values into the channel without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages) // receive from messages
	fmt.Println(<-messages)

}

/*
GoChannel :

a. What is Channel?
Channels are the pipes that connect concurrent goroutines.
Channels allow us to pass data (send & receive) between go routines with the channel operator, <-

b. We can create channel based on:-
	i. Create channel (without buffer)
		_element_ := make(chan _type_)

		- By default, sends and receives block wait until the other side is ready.
		- This allows goroutines to synchronize without explicit locks or condition variables.

	ii. Create channel with buffer
		_element_ := make(chan _type_, _buffer-length_)

		- Sends to a buffered channel block only when the buffer is full.
		- Receives block when the buffer is empty.

Extra Note:
i. Send & Receive data (using operator <-)

	ch <- v
		- Send v to channel ch.

	v := <-ch
		- Receive from ch, and assign value to v.

ii. Beware of deadlocks!!
	How it happens:-
	-  end the buffer without letting the code a chance to read/remove a value from the channel.

*/
