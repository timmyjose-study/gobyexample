package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mayPanic() {
	if rand.Intn(2) == 0 {
		panic("Some error")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// recover is always used inside a deferred function, and control will
	// flow here in case of a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic, error:", r)
		}
	}()

	mayPanic()

	// this line will be reached only if there was no panic in this scope
	fmt.Println("Adios from main!")
}