package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var problem int
	var mu sync.Mutex
	
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				mu.Lock()
				problem++
				mu.Unlock()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(problem)
}
