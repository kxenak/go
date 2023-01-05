package main

import "fmt"

func main() {
	// Basic if-else example.
	if 4%2 == 0 {
		fmt.Println("4 is even.")
	} else {
		fmt.Println("4 is odd.")
	}

	// If without else.
	if 0 == 0 {
		fmt.Println("Of course 0 is equal to 0.")
	}

	// If-else with a statement preceding contitionals.
	if x := 1; x < 0 {
		fmt.Println("x is negative.")
	} else if x < 10 {
		fmt.Println("x has a single digit.")
	} else {
		fmt.Println("x has more than 2 digits.")
	}
}
