package main

import "fmt"

// iota is auto-numbering
const (
	a = iota
	b = iota
	c = iota
)

// after the first one you can leave out iota - cool
const (
	d = iota
	e
	f
)

// and eat the ZERO value
const (
	_ = iota
	g
	h
	i
)

// start at an arbitrary place
const (
	j = iota + 1000
	k
	l
)

const (
	flag1 = 1 << iota
	flag2
	flag3
	flag4
)

func main() {
	fmt.Println("Constants!")

	fmt.Printf("%v, %v, %v\n", a, b, c)
	fmt.Printf("%v, %v, %v\n", d, e, f)
	fmt.Printf("%v, %v, %v\n", g, h, i)
	fmt.Printf("%v, %v, %v\n", j, k, l)

	fmt.Printf("%06b, %06b, %06b, %06b\n", flag1, flag2, flag3, flag4)

	// Lowercase start to avoid exporting it..
	const myConst int = 42

	fmt.Printf("%v, %T\n", myConst, myConst)

	// constants without a type work like #define in c..
	const dave = 42

	var a int16 = 10
	var b int32 = 100

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", a+dave, a+dave)
	fmt.Printf("%v, %T\n", b+dave, b+dave)
}
