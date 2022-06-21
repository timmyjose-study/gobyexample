package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	panicFunction()
}

func panicFunction() {
	randVal := rand.Intn(3)
	if randVal == 1 {
		panic("Got a 1!")
	} else {
		fmt.Println("Got a", randVal)
	}
}