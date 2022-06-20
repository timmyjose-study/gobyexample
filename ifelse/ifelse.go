package main

import "fmt"

func readNum() int {
	var n int
	fmt.Scanf("%d", &n)
	return n
}

func myAbs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func main() {
	n := readNum()
	fmt.Printf("Absolute value of %d = %d\n", n, myAbs(n))
}