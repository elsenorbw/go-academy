package main

import "fmt"

func main() {
	fmt.Println("Functions!")
	for i := 0; i < 5; i++ {
		sayMessage("Hello", i)
	}

	a := sum(1, 2, 3, 4, 5)
	fmt.Println(a)

	b := sum2(1, 2, 3, 4, 6, 7, 7)
	fmt.Println(b)

	d, err := divide(200, 12)
	if err != nil {
		fmt.Printf("Oh no - %v\n", err)
	} else {
		fmt.Printf("Division yields %v\n", d)
	}

	e, err := divide(200, 0)
	if err != nil {
		fmt.Printf("Oh no - %v\n", err)
	} else {
		fmt.Printf("Division yields %v\n", e)
	}

	f := divide
	fmt.Printf("Func is %v, %T\n", f, f)

	greetingTest()
}

func sayMessage(msg string, idx int) {
	fmt.Printf("(%v) %s\n", idx, msg)
}

// handling multiple variables..
func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Printf("The sum is %v\n", result)
	return result
}

// horrible early declare of result variable..
func sum2(values ...int) (result int) {
	for _, v := range values {
		result += v
	}

	return // no need to specify what we are returning.. also weird.
}

// int types for both a and b, syntax sugar
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by zero you fool")
	}
	return a / b, nil
}

// Methods - functions that belong to structs (or whatever)
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Printf("%v %v\n", g.greeting, g.name)
	// g is a copy in this case, any maniplation will be discarded..
	g.name = "FAIL FACE"
}

func (g *greeter) greet2() {
	fmt.Printf("%v %v\n", g.greeting, g.name)
	// g is a pointer, any maniplation will affect the original object..
	g.name = "WORK FACE"
}

func greetingTest() {
	g := greeter{
		greeting: "Hello",
		name:     "Steve",
	}
	g.greet()
	fmt.Printf("Greeting after greet() %v\n", g)
	g.greet2()
	fmt.Printf("Greeting after greet2() %v\n", g)

}
