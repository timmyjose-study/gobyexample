package main

import (
	"fmt"
	"sort"
)

type byLength []string

// implement the Sort interface

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func main() {
	fruits := []string{"apple", "banana", "orange", "lychee", "muskmelon", "pineapple"}
	sort.Strings(fruits)
	fmt.Println("Sorting fruits in lexicographical order, fruits:", fruits)
	sort.Sort(byLength(fruits))
	fmt.Println("Sorting fruits in order of length of name, fruits:", fruits)

}