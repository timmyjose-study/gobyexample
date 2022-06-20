package main

import (
	"fmt"
	"math"
)

const s string = "This is a constant string"

func main() {
	fmt.Println(s)

	const d = 12345
	fmt.Println(d)
	fmt.Println(int64(d))

	const f = 1e10 / 3.9
	fmt.Println(f)

	fmt.Println(math.Sin(f))

}