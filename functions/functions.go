package main

import "fmt"

func plus(n int, m int) int {
	return n + m
}

func plusPlus(a, b, c int) int {
	return plus(a, b) + c
}

func getNum() int {
	var n int
	fmt.Scanf("%d", &n)
	return n
}

func main() {
	a, b, c := getNum(), getNum(), getNum()
	fmt.Println(plus(a, b))
	fmt.Println(plusPlus(a, b, c))

}