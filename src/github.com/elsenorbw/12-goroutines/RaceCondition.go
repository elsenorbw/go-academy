package main

import (
	"fmt"
	"sync"
)

// If we run this with -race then it will generate a warning about lines 19 and 22

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("GO Routines!")
	wg.Add(1)
	go sayHello()
	msg := "Bonjour"
	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	msg = "Guten Tag"
	// crappy but effective..
	wg.Wait()
	fmt.Println("Main func is over")
}

func sayHello() {
	fmt.Println("Hello!")
	wg.Done()
}
