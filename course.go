package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	n := runtime.NumCPU()
	ch := make(chan int)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		go func(id int) {
			for x := range ch {
				fmt.Println(id, x)
			}
		}(i)
	}
	for j := 0; j < 100; j++ {
		ch <- j
	}
	close(ch)
	wg.Wait()
}
