package main

import "fmt"

func main() {
	// A basic slice.
	a := make([]int, 3)
	fmt.Println(a)
	a[0] = 10
	a[1] = 9
	a[2] = 8
	fmt.Println(a)

	// Append method.
	a = append(a, 7)
	a = append(a, 6, 5)
	fmt.Println(a)

	// Copy method.
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(b)

	// Indexing and slicing.
	fmt.Println(b[1:4])
	fmt.Println(b[:2])
	fmt.Println(b[2:])

	// Declaring and initializing in the same expression.
	t := []string{"a", "b", "c"}
	fmt.Println(t)

	// Multidimensional slices.
	x := make([][]int, 5)
	for i := 0; i < len(x); i++ {
		innerLen := i + 1
		x[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			x[i][j] = i - j
		}
	}
	fmt.Println(x)
}
