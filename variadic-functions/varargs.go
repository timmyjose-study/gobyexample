package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println("Got nums:", nums)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3, 4, 5))

	vals := []int{11, 12, 13, 14, 145}
	fmt.Println(sum(vals...))

}