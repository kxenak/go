package main

import (
	"fmt"
	"math"
)

const s string = "bye bye"

func main() {
	fmt.Println(s)
	const num = 500000000
	const v = 3e20 / num
	fmt.Println(v)
	fmt.Println(int64(v))
	fmt.Println(math.Sin(num))
}
