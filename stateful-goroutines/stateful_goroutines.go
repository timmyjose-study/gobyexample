package main

import (
	"fmt"
	"sync"
)

func main() {
	const numWorkers = 50
	var counter uint64
	var wg sync.WaitGroup

	results := make(chan uint64, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			sum := uint64(0)
			for i := 0; i < 1000; i++ {
				sum++
			}
			results <- sum
		}()
	}

	wg.Wait()
	close(results)

	for res := range results {
		counter += res
	}

	fmt.Println("Final counter:", counter)
}