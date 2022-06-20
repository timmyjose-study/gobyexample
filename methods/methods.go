package main

import "fmt"

type rect struct {
	width, length int
}

func (r *rect) area() int {
	return r.width * r.length
}

func (r rect) perimeter() int {
	return 2 * (r.width + r.length)
}

func main() {
	r := rect{width: 10, length: 20}
	fmt.Println(r.area(), r.perimeter())

	r2 := &rect{width: 20, length: 30}
	fmt.Println(r2.area(), r2.perimeter())
}