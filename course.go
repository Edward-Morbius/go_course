package main

import (
	"fmt"
	"time"
)

func main() {
	//Capacity of channel is two: ch := make(chan int, 2)
	ch := make(chan int)
	go func() {
		// No index for channel ranges.
		for x := range ch {
			fmt.Println("~~~", x)
		}
	}()
	ch <- 100
	ch <- 101
	ch <- 102
	time.Sleep(time.Second)
}
