package main

import (
	"fmt"
	"runtime"
)


func main() {
	n := runtime.NumCPU()
	ch := make(chan int)
	fmt.Println(n)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for x := range ch {
				fmt.Println(i, x)
			}
		}(i)
	}
	for j := 0; j < 10; j++ {
		ch <- j
	}
}
