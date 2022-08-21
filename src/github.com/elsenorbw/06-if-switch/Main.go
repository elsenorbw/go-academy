package main

import "fmt"

func main() {
	fmt.Println("If and Switch")

	// v simple
	if false {
		fmt.Println("is true.. oh yes.. very nice")
	} else {
		fmt.Println("is false.. the shame.")
	}

	// if with initialiser statements..
	sp := map[string]int{
		"Cali":  2342453,
		"Texas": 4546456,
		"Boba":  3453456,
	}

	if pop, ok := sp["Bob"]; ok {
		fmt.Printf("Bob population is %v\n", pop)
	} else {
		fmt.Println("Bob not found in the map..")
	}

	value := 17

	if value < 1 {
		fmt.Printf("Your value %v is less than 1 and will not be processed\n", value)
	} else if value > 100 {
		fmt.Printf("Your value %v is greater than 100 and will not be processed\n", value)
	} else {
		fmt.Printf("Your value %v is in range, we will process it\n", value)
	}

	// switch (simple)
	switch i := 2 + 5; i {
	case 1, 2, 5:
		fmt.Printf("%v is 1, 2 or five \n", i)
	case 4, 6, 8:
		fmt.Printf("%v is eight, four or six \n", i)
	default:
		fmt.Printf("%v is some other thing\n", i)
	}

	// switch, no value
	j := 18
	switch {
	case j < 10:
		fmt.Printf("%v is less than 10\n", j)
	case j > 100:
		fmt.Printf("%v is more than 100\n", j)
	default:
		fmt.Printf("Not sure why we're talking about %v\n", j)
	}

	// without the tag, cases can overlap
	j = 28
	switch {
	case j < 10:
		fmt.Printf("%v is less than 10\n", j)
	case j < 20:
		fmt.Printf("%v is less than 20\n", j)
	case j < 30:
		fmt.Printf("%v is less than 30\n", j)
	}

	// You can fallthrough on purpose, but the logic for the case does not work
	j = 6
	switch {
	case j < 10:
		fmt.Printf("%v is less than 10\n", j)
		fallthrough
	case j > 200:
		fmt.Printf("%v is more than 200\n", j)
	case j < 30:
		fmt.Printf("%v is less than 30\n", j)
	}

	// And you can switch on type
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case string:
		fmt.Println("i is a string")
	case float64:
		fmt.Println("i is a float64")
	default:
		fmt.Printf("i is something else : %T\n", i)
	}

}
