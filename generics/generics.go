package main

import "fmt"

func mapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	elem T
	next *element[T]
}

func (lst *List[T]) push(e T) {
	if lst.tail == nil {
		lst.tail = &element[T]{elem: e}
		lst.head = lst.tail
	} else {
		lst.tail.next = &element[T]{elem: e}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) getAll() []T {
	var elems []T

	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.elem)
	}
	return elems
}

func main() {
	m := map[string]int{"hello": 1, "world": 2}
	fmt.Println(mapKeys(m))

	list := List[int]{}
	list.push(1)
	list.push(2)
	fmt.Println(list.getAll())
}