package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, length float64
}

func (r rect) area() float64 {
	return r.width * r.length
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.length)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2.0 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("shape:", g)
	fmt.Println(g.area(), g.perimeter())
}

func main() {
	r := rect{width: 10, length: 20}
	measure(r)

	c := circle{radius: 20.124}
	measure(c)

}