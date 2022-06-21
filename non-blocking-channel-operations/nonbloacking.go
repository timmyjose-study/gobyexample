package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received")
	}

	select {
	case messages <- "hello":
		fmt.Println("Sent message")
	default:
		fmt.Println("No messages sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("No activity")
	}
}