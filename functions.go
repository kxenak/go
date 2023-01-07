package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func multiplus(a, b, c int) int {
	return a + b + c
}

func main() {
	what := plus(1, 56)
	fmt.Println(what)
	what2 := multiplus(1, 2, 4)
	fmt.Println(what2)
}
