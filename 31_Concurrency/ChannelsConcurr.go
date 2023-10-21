/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
)
// Channels are a fundamental synchronization and communication mechanism in Go.
// They allow goroutines to safely communicate and synchronize with each other.
// Channels are typed conduits through which you can send and receive values.
//
// In this example, we create a channel 'channel' to communicate strings.
// The 'message' function sends a message into the channel using the "<-" operator.
// The 'main' function starts a goroutine that sends a message and then receives it.
// The received message is printed, demonstrating how channels enable communication
// between concurrently executing goroutines.

func message(channel chan string) {
	channel <- "This is a message"
	close(channel)
}

func main() {
	channel := make(chan string)

	go message(channel)

	b := <-channel
	fmt.Println(b)
	fmt.Println("Done")
}
