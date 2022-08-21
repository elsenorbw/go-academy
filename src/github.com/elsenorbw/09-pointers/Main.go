package main

import "fmt"

func simple() {
	a := 42
	b := &a
	var c *int = &a
	fmt.Printf("a=%v, b=%v, c=%v\n", a, *b, c)
	a = 27
	fmt.Printf("a=%v, b=%v, c=%v\n", a, *b, c)
}

func pointerMagic() {
	var ms *myStruct
	ms = new(myStruct)
	// Magic compiler help means you can use . instead of -> (or instead of (*p). )
	ms.foo = 123
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}

func main() {
	fmt.Println("Pointers!")
	simple()
	pointerMagic()
}
