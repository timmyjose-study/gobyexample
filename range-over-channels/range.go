package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for j := range jobs {
			fmt.Println("Received job", j)
		}
		done <- true
		fmt.Println("All jobs received")
	}()

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("Sent job", i)
	}

	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
}