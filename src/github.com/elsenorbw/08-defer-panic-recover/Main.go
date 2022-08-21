package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sequences() {
	// sequencing..
	fmt.Println("start")
	// defer pushes this to after the end of the function, lifo
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func deferClose() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", robots)
}

func variablesForDeferHappenAtDeferTime() {
	a := 1
	defer fmt.Println(a)
	a = 7002
}

func DivideByZeroPanic() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}

func MyOwnPanic() {
	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// If we don't want to recover at this point then we need to re-panic
			// panic(err)
		}
	}()
	if true {
		panic("There's something nasty in the Woodshead!")
	}
	fmt.Println("End")
}

func main() {
	fmt.Println("Defer, panic, recover!")
	variablesForDeferHappenAtDeferTime()
	MyOwnPanic()
	fmt.Println("End of Defer, panic, recover!")
}
