package main

import "fmt"

func main() {
	// Basic array.
	var a [5]int
	fmt.Println(a)

	// Indexing an array.
	a[2] = 69
	fmt.Println(a[2], a[3])

	// Lenght of an array.
	fmt.Println("Length of array a: ", len(a))

	// Declaring and initializing together.
	b := [5]int{4, 2, 5, 7, 1}
	fmt.Println(b)

	// Two-dimensional array.
	var c [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = 2*i + 2*j
		}
	}
	fmt.Println(c)
}
