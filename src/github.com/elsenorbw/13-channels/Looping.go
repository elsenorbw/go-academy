package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Channels!")

	// Create a channel for sending and receiving ints
	// 50 is the size of the buffer..
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		// Get data from the channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		// Send data into the channel
		ch <- 42
		// Send a second thing which won't be processed, this will deadlock on an unbuffered channel
		ch <- 27
		// And we need to close the channel
		close(ch)
		wg.Done()
	}(ch)

	// and wait for the routines to finish
	wg.Wait()
	fmt.Println("Main is done")
}
