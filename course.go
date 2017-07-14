package main

import (
	"fmt"
	"sync"
)

func main() {
	var problem int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				problem++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(problem)
}
