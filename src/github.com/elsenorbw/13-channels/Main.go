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

	go func() {
		// Get data from the channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		// Send data into the channel
		ch <- 42
		wg.Done()
	}()

	// and wait for the routines to finish
	wg.Wait()
	fmt.Println("Main is done")
}
