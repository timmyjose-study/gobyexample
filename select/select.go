package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		c2 <- "two"
	}()

loop:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout")
			break loop

		}
	}
}