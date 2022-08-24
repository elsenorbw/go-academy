package main

import (
	"fmt"

	"github.com/elsenorbw/15-modules/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
