package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func main() {
	n := 42
	fmt.Println(n)

	zeroval(n)
	fmt.Println(n)

	zeroptr(&n)
	fmt.Println(n)
	fmt.Println(&n)
}