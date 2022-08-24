package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("GO Routines!")
	wg.Add(1)
	go sayHello()
	msg := "Bonjour"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Guten Tag"
	// crappy but effective..
	wg.Wait()
	fmt.Println("Main func is over")
}

func sayHello() {
	fmt.Println("Hello!")
	wg.Done()
}
