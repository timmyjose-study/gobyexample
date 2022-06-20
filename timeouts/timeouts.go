package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

	select {
	case res1 := <-c1:
		fmt.Println(res1)
	case <-time.After(time.Second):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		c2 <- "Result 2"
	}()

	select {
	case res2 := <-c2:
		fmt.Println(res2)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout 2")
	}

}