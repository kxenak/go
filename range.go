package main

import "fmt"

func main() {
	// Basic example of range.
	a := []string{"hello", "bye", "hi", "go"}
	for _, val := range a {
		fmt.Println(val)
	}

	// Range with indexing.
	b := []string{"hello", "bye", "hi", "go"}
	for i, val := range b {
		fmt.Println(i, val)
	}

	// Range with maps.
	m := map[string]string{"foo": "boo", "boo": "loo"}
	for key, val := range m {
		fmt.Println(key, val)
	}

	// Range with strings.
	for x, y := range "hello" {
		fmt.Println(x, y)
	}
}
