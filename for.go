package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i + 13)
		i++
	}
	for i := 0; i < 3; i++ {
		fmt.Println(i + 39)
	}
	for {
		fmt.Println("loop loop")
		break
	}
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
