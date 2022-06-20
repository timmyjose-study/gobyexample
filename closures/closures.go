package main

import "fmt"

func linum() func() {
	line := 0
	return func() {
		line++
		fmt.Println("Current line:", line)
	}
}

func main() {
	l := linum()
	for i := 0; i < 10; i++ {
		l()
	}

	l = linum()

	for i := 0; i < 5; i++ {
		l()
	}
}