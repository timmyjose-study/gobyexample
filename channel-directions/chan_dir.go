package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func send(s <-chan string, r chan<- string) {
	defer wg.Done()
	msg := <-s
	r <- msg
}

func receive(c <-chan string) {
	defer wg.Done()
	fmt.Println(<-c)
}

func main() {
	wg.Add(2)

	sender := make(chan string, 1)
	receiver := make(chan string, 1)

	go send(sender, receiver)
	sender <- "Some message"
	go receive(receiver)

	wg.Wait()
}