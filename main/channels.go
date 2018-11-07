package main

import (
	"fmt"
	"time"
)

func displayThroughChannel(c <-chan int) { // <- is optional, in this case it indicates we cannot send data through this channel
	for v := range c {
		fmt.Printf("I've just received %v\n", v)
	}
}

func main() {
	/*
		Go main philosophy:
		Don't communicate by sharing memory, share memory by communicating
	*/

	c := make(chan int, 1024)
	// Buffered channel, 1024 maximum messages can be buffered before to block the sender goroutine (not the thread)
	go displayThroughChannel(c)

	c <- 5
	c <- 7

	close(c) // The best practice is to close a channel from the sender
	time.Sleep(250 * time.Millisecond)

	// c <- 8 // We can't send value to a closed channel, it will panic

	/*
		Benefits of communicating by sharing memory:
		- No lock
		- No false sharing
		- Etc.
	*/

	/*
	   TODO
		Create two functions communicating by channels (req/reply)
	*/
}
