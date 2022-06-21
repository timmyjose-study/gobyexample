package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"hello", "world", "nice", "to", "meet", "you"}
	fmt.Println("Before sorting, strs:", strs)
	sort.Strings(strs)
	fmt.Println("After sorting, strs:", strs)
	fmt.Println(sort.StringsAreSorted(strs))

	ints := []int{5, 1, 2, 4, 3}
	fmt.Println("Before sorting, ints:", ints)
	sort.Ints(ints)
	fmt.Println("After sorting, ints:", ints)

	fmt.Println(sort.IntsAreSorted(ints))
}