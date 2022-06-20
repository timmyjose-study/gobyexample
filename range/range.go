package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)

	for i, n := range nums {
		fmt.Printf("%d: %d\n", i, n)
	}

	kvs := map[string]string{"a": "apple", "b": "ball", "c": "cat"}
	for k, v := range kvs {
		fmt.Printf("%v => %v\n", k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}

	for _, v := range kvs {
		fmt.Println(v)
	}
}