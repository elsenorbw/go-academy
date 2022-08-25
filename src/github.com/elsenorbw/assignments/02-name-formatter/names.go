package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.Stdin is of type %T\n", os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your first name:")
	first, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v\n", first)
}
