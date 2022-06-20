package main

import (
	"fmt"
	"time"
)

func worker(c chan<- bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")
	c <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	<-done
}