package main

import "fmt"

// ClosingChannel :
// Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.
func ClosingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Receiver
	go func() {
		// It repeatedly receives from jobs with j, more := <-jobs. In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received. We use this to notify on done when we’ve worked all our jobs.
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Sender
	for j := 1; j <= 3; j++ {
		jobs <- j // send job(sender) to worker (receiver)
		fmt.Println("sent job", j)
	}
	close(jobs) // Closing a channel indicates that no more values will be sent on it
	fmt.Println("sent all jobs")

	<-done // We await the worker using the synchronization approach we saw earlier.
}

/*
i. Sender
	- A sender can close a channel to indicate that no more values will be sent.
ii. Receiver
	- Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression.
iii. Expression
	- v, ok := <-ch , where v as a sender, ok as a receiver. (ok is false if there are no more values to receive and the channel is closed.)

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren’t like files; you don’t usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
*/
