package main

import "fmt"

func vals() (int, int) {
	return 4, 11
}

func main() {
	a, b := vals()
	fmt.Printf("%d + %d = %d\n", a, b, func(n, m int) int { return n + m }(a, b))
}