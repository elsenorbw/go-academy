package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Channels!")

	// Create a channel for sending and receiving ints
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) {
		// Get data from the channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		// Send data into the channel
		ch <- 42
		wg.Done()
	}(ch)

	// and wait for the routines to finish
	wg.Wait()
	fmt.Println("Main is done")
}
