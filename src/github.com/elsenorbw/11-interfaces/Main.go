package main

import "fmt"

func main() {
	fmt.Println("Interfaces!")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello go !"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// So this implements the interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Need to see about functions with different return types and everything else the same.. feels like it must work
// Apparently this is impossible and the built-in language features that do it are cheating.
