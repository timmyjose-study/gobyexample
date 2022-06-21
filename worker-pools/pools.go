package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "is working on job", j)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Worker", id, "has finished job", j)
		results <- j * 2
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < numJobs; i++ {
		<-results
	}
}