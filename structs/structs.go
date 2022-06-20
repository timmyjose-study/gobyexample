package main

import "fmt"

type person struct {
	name string
	age  int
}

func makePerson(name string) *person {
	p := person{name: name, age: 42}
	return &p
}

func main() {
	fmt.Println(makePerson("Bob"))
	fmt.Println(person{name: "Wanda", age: 99})

	s := makePerson("Sean")
	s.age = 100
	fmt.Println(*s)
}