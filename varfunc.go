// Example of variadic functions that take any number of trailing arguments.

package main

import "fmt"

func sum(num ...int) int {
	total := 0
	for _, val := range num {
		total += val
	}
	return total
}

func main() {
	sum1 := sum(1, 2, 34)
	sum2 := sum(1, 3)
	a := []int{2, 3, 4, 5, 6, 5}
	sum3 := sum(a...)
	fmt.Println(sum1, sum2, sum3)
}
