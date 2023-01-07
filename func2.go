// A function with multiple return values.

package main

import "fmt"

func val() (int, int, int) {
	return 55, 66, 77
}

func main() {
	a, b, c := val()
	fmt.Println(a, b, c)
}
