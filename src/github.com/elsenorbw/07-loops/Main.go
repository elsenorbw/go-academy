package main

import "fmt"

func main() {
	fmt.Println("Loops!")

	for i := 0; i < 5; i++ {
		fmt.Printf("Looping %v\n", i)
	}

	// Looping with multi counters..
	for i, j := 0, 10; i < 5; i, j = i+2, j+1 {
		fmt.Printf("Looping %v, %v\n", i, j)
	}

	// statements are optional
	i := 17
	for ; i < 6; i += 2 {
		fmt.Printf("Loop with no init: %v\n", i)
	}

	// for as a while loop, you can drop both semicolons but not just one..

	for i > 10 {
		// hmm
		fmt.Printf("Loop with no increment.. %v\n", i)
		if i%2 == 0 {
			i -= 3
		} else {
			i += 1
		}
	}

	// continue works like c
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Printf("Odd value %v\n", j)
	}

	// break allows labels to break out to a level :)

flibart:
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b += 2 {
			for c := 0; c < 10; c += 3 {
				TheVal := a * b * c
				fmt.Printf("%v, %v, %v -> %v\n", a, b, c, TheVal)
				if TheVal > 100 {
					break flibart
				}
			}
		}
	}

	// Now that's nice

	// working with collections..
	s := []int{1, 2, 3, 4, 50, 60, 7, 8}
	for k, v := range s {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	// Maps..
	populations := map[string]int{
		"abc": 123,
		"def": 456,
		"ghi": 789,
	}

	for k, v := range populations {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	// Just get the keys
	for k := range populations {
		fmt.Printf("k=%v\n", k)
	}

	// Just get the values
	for _, v := range populations {
		fmt.Printf("v=%v\n", v)
	}

}
