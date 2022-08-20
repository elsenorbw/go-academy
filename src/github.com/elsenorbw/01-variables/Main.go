package main

import "fmt"

// package level variables
var m_i int = 67

// var block
var (
	actor               = "Keanu Reeves"
	character           = "Neo"
	my_unused_variables = 78
)

// package scope var, overridden by the local var..
var i int = 100

func type_conversions() {
	var a int = 42
	var b float32 = float32(a)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
}

func string_functions() {
	// start with an int
	a := 65
	s1 := string(a)
	fmt.Printf("s1 - using string() yields %v, %T\n", s1, s1)
	s2 := fmt.Sprintf("%v", a)
	fmt.Printf("s2 - using Fmt.Sprintf yields %v, %T\n", s2, s2)

}

func main() {

	fmt.Println(i)
	// declaration with type
	var i int
	i = 42

	// declaration with type and initialisation
	var j int = 43

	// type inference with declarations
	k := 44

	fmt.Println(i, j, k)

	fmt.Printf("%v, %T\n", j, j)

	fmt.Println(actor, character)

	type_conversions()
	string_functions()
}
