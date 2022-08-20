package main

import "fmt"

func main() {
	var a bool = true
	fmt.Printf("bool true gives us %v, %T\n", a, a)

	// sized integers
	// int minimum of 32 bits
	n := 129
	fmt.Printf("default int type: %v, %T\n", n, n)

	var m int8 = int8(n)
	fmt.Printf("default int type: %v, %T\n", m, m)

	// complex is a thing, real() and imag() get the bits..
	// complex64 (float32+float32)
	// complex128 (float64+float64)
	// syntax 4+2i
	c := 34 + 2i
	fmt.Printf("complex number %v, %T\n", c, c)

	// strings are immutable (boo, no fun)
	s := "This is a string"
	b := []byte(s)
	fmt.Printf("string %v, %T\n", s, s)
	fmt.Printf("bytes %v, %T\n", b, b)

	r := 'a'
	fmt.Printf("rune %v, %T\n", r, r)

}
