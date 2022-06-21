package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter uint64

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}

		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}