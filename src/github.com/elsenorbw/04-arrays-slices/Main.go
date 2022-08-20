package main

import "fmt"

func main() {
	fmt.Println("Arrrays and Slices!")

	// Explicit sizing..
	grades := [3]int{101, 102, 103}
	fmt.Printf("%v\n", grades)

	// Implicit sizing from initialisation
	dave := [...]int{1, 2, 3, 4}
	fmt.Printf("%v, %T\n", dave, dave)

	// 0 based like God intended
	dave[1] = 9001
	fmt.Printf("%v, %T\n", dave, dave)

	// Arrays know their length
	fmt.Printf("Size of dave %v\n", len(dave))

	// Array copies copy everything...
	ernie := dave
	dave[2] = 8000
	fmt.Printf("dave: %v, %T\n", dave, dave)
	fmt.Printf("ernie: %v, %T\n", ernie, ernie)

	// but there are pointers, hurrah
	fred := &dave
	dave[3] = 7000
	fmt.Printf("dave: %v, %T\n", dave, dave)
	fmt.Printf("ernie: %v, %T\n", ernie, ernie)
	fmt.Printf("fred: %v, %T\n", fred, fred)

	// Slices
	a := []int{1000, 1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010, 1011, 1012}
	fmt.Printf("%v, %T\n", a, a)
	// slices are references to data, so this will point at the same thing
	b := a
	b[2] = 200
	fmt.Printf("a=%v, %T\n", a, a)
	fmt.Printf("b=%v, %T\n", b, b)

	// And it supports Python slicing styles..

	c := a[:]   // everything
	d := a[3:]  // start from index 3
	e := a[:6]  // everything up to but not including index 6
	f := a[3:6] // stuff from index 3 up to but not including index 6

	fmt.Printf("c=%v, %T, len=%v, cap=%v\n", c, c, len(c), cap(c))
	fmt.Printf("d=%v, %T, len=%v, cap=%v\n", d, d, len(d), cap(d))
	fmt.Printf("e=%v, %T, len=%v, cap=%v\n", e, e, len(e), cap(e))
	fmt.Printf("f=%v, %T, len=%v, cap=%v\n", f, f, len(f), cap(f))

	// using make to create a slice..
	z := make([]int, 3, 100)
	fmt.Printf("c=%v, %T, len=%v, cap=%v\n", z, z, len(z), cap(z))

	z = append(z, 100)
	fmt.Printf("c=%v, %T, len=%v, cap=%v\n", z, z, len(z), cap(z))

	z = append(z, 102, 103, 104, 105, 106)
	fmt.Printf("c=%v, %T, len=%v, cap=%v\n", z, z, len(z), cap(z))

	// adding slices together.. you can spread the slice contents..
	var myslice []int = []int{40, 41, 42, 43}
	fmt.Printf("c=%v, %T, len=%v, cap=%v\n", myslice, myslice, len(myslice), cap(myslice))

	// appending that slice with spread
	z = append(z, myslice...)
	fmt.Printf("c=%v, %T, len=%v, cap=%v\n", z, z, len(z), cap(z))

}
